package telegram

import "encoding/json"

// GetMe returns basic information about the bot.
// It's a simple method for testing your bot's auth token.
// GetMe also update bot User info.
//
// https://core.telegram.org/bots/api#getme.
func (bot *Bot) GetMe() (User, error) {
	resp, err := bot.MakeRequest("getMe", nil)
	if err != nil {
		return User{}, err
	}

	var user User
	if err := json.NewDecoder(resp).Decode(&user); err != nil {
		return User{}, err
	}
	bot.User = user

	return bot.User, nil
}

// SendMessage sends text messages. On success, the sent Message is returned.
//
//  TODO: Params: SetParseMode, SetEntities, SetDisableWebPagePreview, SetDisableNotification, SetReplyToMessageID, SetAllowSendingWithoutReply, SetReplyMarkup
//
// https://core.telegram.org/bots/api#sendmessage
func (bot *Bot) SendMessage(chatID int, text string, params ...Param) (Message, error) {
	params = append(params, setParamInt("chat_id", chatID), setParamString("text", text))
	urlVal := resolveParam(params)
	resp, err := bot.MakeRequest("sendMessage", urlVal)
	if err != nil {
		return Message{}, err
	}

	var message Message
	if err := json.NewDecoder(resp).Decode(&message); err != nil {
		return Message{}, err
	}

	return message, nil
}
