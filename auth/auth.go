package auth

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var Bot_settings = tb.Settings{
	Token: "xxxxxxxxxxxx",
	// You can also set custom API URL. If field is empty it equals to "https://api.telegram.org"
	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
}

func GetConnection() (*tb.Bot, error) {
	return tb.NewBot(Bot_settings)
}
