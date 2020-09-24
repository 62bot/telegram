package telegram_test

import (
	"testing"

	"github.com/botcare/telegram"
)

func TestHelloWorld(t *testing.T) {
	if telegram.Greet() != "Hello World! I'm Telegram Bots" {
		t.Error("Telegram Bots is not responding")
	}
}
