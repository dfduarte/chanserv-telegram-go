package hello

import (
	"github.com/dfduarte/chanserv-telegram-go/auth"
	tb "gopkg.in/tucnak/telebot.v2"
)

func Hello() {

	var bot, _ = auth.GetConnection()

	bot.Handle("/hello", func(m *tb.Message) {
		bot.Send(m.Sender, "hello world")
	})

	bot.Handle(tb.OnUserJoined, func(m *tb.Message) {
		bot.Send(m.Sender, "Howdy User")
	})

	bot.Start()
}
