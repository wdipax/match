package adapter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/state"
)

type Adapter struct {
	bot     *tgbotapi.BotAPI
	isAdmin func(*tgbotapi.User) bool

	state *state.State
}

func New(bot *tgbotapi.BotAPI, isAdmin func(*tgbotapi.User) bool) *Adapter {
	s := state.New()

	a := Adapter{
		bot:     bot,
		isAdmin: isAdmin,
		state:   s,
	}

	return &a
}

func (a *Adapter) Process(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	e := event.Event{
		ChatID:    update.FromChat().ChatConfig().ChatID,
		FromAdmin: a.isAdmin(update.SentFrom()),
	}

	r := a.state.Process(&e)

	// keyboard1 := tgbotapi.NewOneTimeReplyKeyboard(
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("start session"),
	// 	),
	// )

	for _, m := range r.GetMessages() {
		msg := tgbotapi.NewMessage(m.ChatID, m.Text)

		a.bot.Send(msg)
	}
}
