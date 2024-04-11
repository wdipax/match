package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/adapter"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	a := adapter.New(bot, update)

	// TODO: shutdown on receiving termination signal.
	for update := range updates {
		// TODO: is it benefitial to use the sync pool here?
		// TODO: skip all messages created before the bot has started.
		// TODO: serve concurrently.
		a.Process(update)
	}
}
