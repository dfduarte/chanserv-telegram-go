package hello

import (
	"fmt"

	"github.com/dfduarte/chanserv-telegram-go/auth"
	tb "gopkg.in/tucnak/telebot.v2"
)

var userInput string
var fmtUser string

func joinedUsername(userInput string) (fmtUser string) {
	fmtUser = fmt.Sprintf("Olá usuário %s, seja bem vindo!", userInput)
	return
}

func Hello() {

	var bot, _ = auth.GetConnection()

	bot.Handle(tb.OnUserJoined, func(m *tb.Message) {
		outputMessage := joinedUsername(m.Sender.FirstName)
		bot.Send(m.Chat, outputMessage)

	})

	/*
		bot.Handle(tb.OnUserJoined, func(m *tb.Message) {
			bot.Send(m.Sender, "Howdy User")
		})
	*/

	bot.Start()
}
