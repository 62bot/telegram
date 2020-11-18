package telegram

import "net/http"

// TelegramURL is a Telegram Host URL.
const TelegramURL = "https://api.telegram.org"

// The Option is used for tweaking the bot settings.
type Option func(*Bot)

// SetHostURL returns an option to set bot host URL.
func SetHostURL(hostURL string) Option {
	return func(bot *Bot) {
		bot.hostURL = hostURL
	}
}

// SetClient returns an option to set bot client.
func SetClient(c *http.Client) Option {
	return func(bot *Bot) {
		bot.client = c
	}
}

// trySetDefaultHostURL sets bot host URL to default if its unset.
func (bot *Bot) trySetDefaultHostURL() {
	if bot.hostURL == "" {
		SetHostURL(TelegramURL)(bot)
	}
}

// trySetDefaultClient sets bot client to default if its unset.
func (bot *Bot) trySetDefaultClient() {
	if bot.client == nil {
		SetClient(http.DefaultClient)(bot)
	}
}
