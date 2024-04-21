package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("new session", "/new_session"),
		),
	)

	// TODO: shutdown on receiving termination signal.
	// TODO: skip all messages created before the bot has started.
	// TODO: serve concurrently.
	for update := range updates {
		if update.Message != nil { // If we got a message

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, " ")
			msg.ReplyMarkup = keyboard

			bot.Send(msg)
		} else if update.CallbackQuery != nil {
			log.Println(update.CallbackData())

			msg := update.CallbackQuery.Message

			del := tgbotapi.NewDeleteMessage(msg.Chat.ID, msg.MessageID)

			bot.Send(del)
		}
	}
}
