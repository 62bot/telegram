package telegram

import "encoding/json"

// Update represents an incoming update.
// At most one of the optional parameters can be present in any given update.
//
// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`              // Optional.
	EditedMessage      *Message            `json:"edited_message,omitempty"`       // Optional.
	ChannelPost        *Message            `json:"channel_post,omitempty"`         // Optional.
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`  // Optional.
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`         // Optional.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"` // Optional.
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`       // Optional.
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`       // Optional.
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`   // Optional.
	Poll               *Poll               `json:"poll,omitempty"`                 // Optional.
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`          // Optional.
}

// WebhookInfo contains information about the current status of a webhook.
//
// https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	URL                  string   `json:"url"`
	HasCustomCertificate bool     `json:"has_custom_certificate"`
	PendingUpdateCount   int      `json:"pending_update_count"`
	IPAddress            string   `json:"ip_address,omitempty"`         // Optional.
	LastErrorDate        int      `json:"last_error_date,omitempty"`    // Optional.
	LastErrorMessage     string   `json:"last_error_message,omitempty"` // Optional.
	MaxConnections       int      `json:"max_connections,omitempty"`    // Optional.
	AllowedUpdates       []string `json:"allowed_updates,omitempty"`    // Optional.
}

// GetUpdates returns a slice of Update.
// Use this method to receive incoming updates using long polling.
//
//  Params: SetOffset, SetLimit, SetTimeout, SetAllowedUpdates.
//
// https://core.telegram.org/bots/api#getupdates
func (bot *Bot) GetUpdates(params ...Param) ([]Update, error) {
	urlVal := resolveParam(params)
	resp, err := bot.MakeRequest("getUpdates", urlVal)
	if err != nil {
		return nil, err
	}

	var updates []Update
	if err := json.NewDecoder(resp).Decode(&updates); err != nil {
		return nil, err
	}

	return updates, nil
}

// SetWebhook specify a url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update.
// In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
//
//  Params: SetCertificate (TODO), SetIPAddress, SetMaxConnections, SetAllowedUpdates, SetDropPendingUpdates.
//
// https://core.telegram.org/bots/api#setwebhook
func (bot *Bot) SetWebhook(url string, params ...Param) (bool, error) {
	params = append(params, setParamString("url", url))
	urlVal := resolveParam(params)
	resp, err := bot.MakeRequest("setWebhook", urlVal)
	if err != nil {
		return false, err
	}

	var ok bool
	if err := json.NewDecoder(resp).Decode(&ok); err != nil {
		return false, err
	}

	return ok, nil
}

// DeleteWebhook removes webhook integration if you decide to switch back to GetUpdates. Returns True on success.
//
//  Params: SetDropPendingUpdates.
//
// https://core.telegram.org/bots/api#deletewebhook
func (bot *Bot) DeleteWebhook(params ...Param) (bool, error) {
	urlVal := resolveParam(params)
	resp, err := bot.MakeRequest("deleteWebhook", urlVal)
	if err != nil {
		return false, err
	}

	var ok bool
	if err := json.NewDecoder(resp).Decode(&ok); err != nil {
		return false, err
	}

	return ok, nil
}

// GetWebhookInfo returns current webhook status. Requires no parameters.
// On success, returns a WebhookInfo object. If the bot is using GetUpdates, will return an object with the url field empty.
//
// https://core.telegram.org/bots/api#getwebhookinfo
func (bot *Bot) GetWebhookInfo() (WebhookInfo, error) {
	resp, err := bot.MakeRequest("getWebhookInfo", nil)
	if err != nil {
		return WebhookInfo{}, err
	}

	var webhookInfo WebhookInfo
	if err := json.NewDecoder(resp).Decode(&webhookInfo); err != nil {
		return WebhookInfo{}, err
	}

	return webhookInfo, nil
}
