package auth

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var botSettings = tb.Settings{
	Token: "649212493:AAEvogFc9QKKUOR7tEXWfyxn9V4cI7Qm2XE",
	// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
}

func GetConnection() (*tb.Bot, error) {
	return tb.NewBot(botSettings)
}
