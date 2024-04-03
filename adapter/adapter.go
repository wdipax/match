package adapter

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Update struct {
	bot    *tgbotapi.BotAPI
	update tgbotapi.Update
}

func New(bot *tgbotapi.BotAPI, update tgbotapi.Update) *Update {
	return &Update{
		update: update,
		bot:    bot,
	}
}

func (u *Update) FromAdmin() bool {
	update := u.update

	if update.Message == nil || update.Message.From == nil {
		return false
	}

	return update.Message.From.ID == 131381334
}

func (u *Update) SendMessage(text string) {
	update := u.update
	bot := u.bot

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
