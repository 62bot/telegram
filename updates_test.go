package telegram_test

import (
	"net/http"

	"github.com/62bot/telegram"
	"github.com/billyzaelani/is"
)

func getUpdatesOK(is *is.Is, bot *telegram.Bot) {
	var (
		wantUpdateLen = 4
		wantFromID    = 12345
		wantChatID    = 12345
	)

	updates, err := bot.GetUpdates()
	is.NoError(err)

	is.Equal(len(updates), wantUpdateLen)
	for i, update := range updates {
		is.Equal(update.UpdateID, i)                 // update_id
		is.Equal(update.Message.MessageID, i)        // message.message_id
		is.Equal(update.Message.From.ID, wantFromID) // message.from.id
		is.Equal(update.Message.Chat.ID, wantChatID) // message.chat.id
	}
}

func getUpdatesWithParams(is *is.Is, bot *telegram.Bot) {
	var (
		offset  = 1
		limit   = 2
		timeout = 3

		wantUpdateLen = limit
		wantFromID    = 12345
		wantChatID    = 12345
	)

	updates, err := bot.GetUpdates(
		telegram.SetOffset(offset),
		telegram.SetLimit(limit),
		telegram.SetTimeout(timeout),
		telegram.SetAllowedUpdates(),
	)
	is.NoError(err)

	is.Equal(len(updates), wantUpdateLen)
	for i, update := range updates {
		is.Equal(update.UpdateID, i+offset)          // update_id
		is.Equal(update.Message.MessageID, i+offset) // message.message_id
		is.Equal(update.Message.From.ID, wantFromID) // message.from.id
		is.Equal(update.Message.Chat.ID, wantChatID) // message.chat.id
	}
}

func getUpdatesConflictWithWebhook(is *is.Is, bot *telegram.Bot) {
	_, err := bot.GetUpdates()
	var botError *telegram.BotError
	is.ErrorAs(err, &botError)
	is.Equal(botError.Code, http.StatusConflict)
}

func setWebhookOK(is *is.Is, bot *telegram.Bot) {
	ok, err := bot.SetWebhook("https://example.com")
	is.NoError(err)

	is.True(ok) // setWebhook should be ok
}

// TODO: func setWebhookWithCertificate(is *is.Is, bot *telegram.Bot) {}

func setWebhookWithParams(is *is.Is, bot *telegram.Bot) {
	ok, err := bot.SetWebhook("https://example.com",
		telegram.SetIPAddress("127.0.0.1"),
		telegram.SetMaxConnections(100),
		telegram.SetAllowedUpdates(),
		telegram.SetDropPendingUpdates(true),
	)
	is.NoError(err)

	is.True(ok)
}

func setWebhookErrorNotHTTPS(is *is.Is, bot *telegram.Bot) {
	_, err := bot.SetWebhook("http://example.com")
	var botError *telegram.BotError
	is.ErrorAs(err, &botError)
	is.Equal(botError.Code, http.StatusBadRequest)
}

func deleteWebhookOK(is *is.Is, bot *telegram.Bot) {
	ok, err := bot.DeleteWebhook(
		telegram.SetDropPendingUpdates(true),
	)
	is.NoError(err)

	is.True(ok)
}

func getWebhookInfoOK(is *is.Is, bot *telegram.Bot) {
	webhookInfo, err := bot.GetWebhookInfo()
	is.NoError(err)

	is.Equal(webhookInfo.URL, "https://example.com")
}
