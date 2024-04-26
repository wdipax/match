package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("creating bot: %s", err)
	}

	// TODO: do we need this?
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	// TODO: shutdown on receiving termination signal.
	// TODO: skip all messages created before the bot has started.
	// TODO: serve concurrently.
	for update := range updates {
		_ = update
		// a.Process(update)
	}
}
