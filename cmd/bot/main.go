package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/nikandpro/notes-bot/pkg/telegram"
)

func main() {

	fmt.Println("Start")

	bot, err := tgbotapi.NewBotAPI("Token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)

	if err = telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
