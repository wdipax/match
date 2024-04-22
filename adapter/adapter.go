package adapter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/response"
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

	e := event.New(
		update.FromChat().ChatConfig().ChatID,
		update.Message.Text,
		a.isAdmin(update.SentFrom()),
		func() string {
			if update.Message.Command() == "start" {
				return update.Message.CommandArguments()
			}
			return ""
		}(),
	)

	r := a.state.Process(e)

	// keyboard1 := tgbotapi.NewOneTimeReplyKeyboard(
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("start session"),
	// 	),
	// )

	for _, m := range r.GetMessages() {
		var text string

		switch m.Type {
		case response.Text:
			text = m.Text
		case response.BoysLink:
			text = "To join boys team click https://t.me/" + a.bot.Self.UserName + "?start=" + m.Text
		case response.GirlsLink:
			text = "To join girls team click https://t.me/" + a.bot.Self.UserName + "?start=" + m.Text
		}

		msg := tgbotapi.NewMessage(m.ChatID, text)

		a.bot.Send(msg)
	}
}
