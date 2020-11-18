package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Bot is a Telegram bot.
//
// https://core.telegram.org/bots
type Bot struct {
	token   string
	hostURL string
	client  *http.Client

	User User // Bot info.
}

// NewBot returns a new bot given a token and other optional options.
func NewBot(token string, options ...Option) (*Bot, error) {
	bot := &Bot{token: token}
	for _, option := range options {
		option(bot)
	}

	bot.trySetDefaultHostURL()
	bot.trySetDefaultClient()

	_, err := bot.GetMe()

	return bot, err
}

// Response represents a general response from Telegram.
//
// https://core.telegram.org/bots/api#making-requests
type Response struct {
	OK          bool                `json:"ok"`
	Description string              `json:"description,omitempty"` // Optional.
	Result      json.RawMessage     `json:"result,omitempty"`      // Optional.
	ErrorCode   int                 `json:"error_code,omitempty"`  // Optional.
	Parameters  *ResponseParameters `json:"parameters,omitempty"`  // Optional.

	buf *bytes.Reader
}

func (r *Response) Read(b []byte) (n int, err error) {
	if r.buf == nil {
		r.buf = bytes.NewReader(r.Result)
	}
	return r.buf.Read(b)
}

// MakeRequest is a raw method to make a request to the Telegram API.
// Generally other bot method like GetMe or GetUpdates is used instead
// unless there's a new method in Telegram API that doesn't handle yet by the bot.
//
// All request are passed to Telegram Bot API in the form:
//  https://api.telegram.org/bot<token>/METHOD_NAME
// Like this for example:
//  https://api.telegram.org/bot123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11/getMe
//
// https://core.telegram.org/bots/api#making-requests
func (bot *Bot) MakeRequest(methodName string, params url.Values) (*Response, error) {
	endpoint := fmt.Sprintf("%s/bot%s/%s", bot.hostURL, bot.token, methodName)

	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	w, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer w.Body.Close()

	var resp Response
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		return nil, err
	}

	if !resp.OK {
		return nil, &BotError{Code: resp.ErrorCode, Description: resp.Description}
	}

	return &resp, nil
}

// BotError records an error code (http status code) and the description.
type BotError struct {
	Code        int
	Description string
}

func (e *BotError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Description)
}
