package adapter

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Adapter struct {
	bot *tgbotapi.BotAPI
}

func New(bot *tgbotapi.BotAPI) *Adapter {
	return &Adapter{
		bot: bot,
	}
}

func (a *Adapter) Process(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	
	keyboard1 := tgbotapi.NewOneTimeReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("start session"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "actions")
	msg.ReplyMarkup = keyboard1

	a.bot.Send(msg)
}
