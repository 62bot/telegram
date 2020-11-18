package telegram_test

import (
	"github.com/62bot/telegram"
	"github.com/billyzaelani/is"
)

type methodTestFunc func(is *is.Is, bot *telegram.Bot)

func sendMessageOK(is *is.Is, bot *telegram.Bot) {
	var (
		wantFromID = 54321
		wantChatID = 12345
		wantText   = "this message will destroy itself in 3...2...1..."
	)

	message, err := bot.SendMessage(wantChatID, wantText)
	is.NoError(err)

	is.Equal(message.MessageID, 0)
	is.Equal(message.From.ID, wantFromID)
	is.Equal(message.Chat.ID, wantChatID)
	is.Equal(message.Text, wantText)
}
