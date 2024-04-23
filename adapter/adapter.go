package adapter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/response"
	"github.com/wdipax/match/state"
)

const (
	stat                   = "show statistics"
	endTeamRegistration    = "end teams registration"
	reopenTeamRegistration = "back to teams registration"
	startVoting            = "start voting"
)

type Adapter struct {
	bot     *tgbotapi.BotAPI
	isAdmin func(*tgbotapi.User) bool

	state *state.State

	control          *control
	teamRegistration *control
	knowEachOther    *control
}

func New(bot *tgbotapi.BotAPI, isAdmin func(*tgbotapi.User) bool) *Adapter {
	s := state.New()

	a := Adapter{
		bot:     bot,
		isAdmin: isAdmin,

		state: s,
	}

	a.teamRegistration = &control{
		keyboard: tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(stat),
				tgbotapi.NewKeyboardButton(endTeamRegistration),
			),
		),
		nextStage: endTeamRegistration,
	}

	a.knowEachOther = &control{
		keyboard: tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton(reopenTeamRegistration),
				tgbotapi.NewKeyboardButton(startVoting),
			),
		),
		previousStage: reopenTeamRegistration,
		nextStage:     startVoting,
	}

	return &a
}

type control struct {
	keyboard      tgbotapi.ReplyKeyboardMarkup
	previousStage string
	nextStage     string
}

func (c *control) previousStageText() string {
	if c == nil {
		return ""
	}

	return c.previousStage
}

func (c *control) nextStageText() string {
	if c == nil {
		return ""
	}

	return c.nextStage
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
		func() int {
			switch update.Message.Text {
			case stat:
				return event.Statistics
			case a.control.previousStageText():
				return event.PreviousStage
			case a.control.nextStageText():
				return event.NextStage
			default:
				return event.Unknown
			}
		}(),
	)

	r := a.state.Process(e)

	for _, m := range r.GetMessages() {
		text := m.Text

		msg := tgbotapi.NewMessage(m.ChatID, text)

		switch m.Type {
		case response.Text:
			text = m.Text
		case response.BoysLink:
			text = "To join boys team click https://t.me/" + a.bot.Self.UserName + "?start=" + m.Text
		case response.GirlsLink:
			text = "To join girls team click https://t.me/" + a.bot.Self.UserName + "?start=" + m.Text
		case response.TeamRegistration:
			a.control = a.teamRegistration

			msg.ReplyMarkup = a.control.keyboard
		case response.KnowEachOther:
			a.control = a.knowEachOther

			msg.ReplyMarkup = a.control.keyboard
		}

		msg.Text = text

		if a.control == nil {
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		a.bot.Send(msg)
	}
}
