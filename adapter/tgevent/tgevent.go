package tgevent

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/protocol/command"
	"github.com/wdipax/match/protocol/step"
)

type TGEvent struct {
	tgbotapi.Update
	admin string
	stage int
}

func New(update tgbotapi.Update, admin string, stage int) *TGEvent {
	return &TGEvent{
		Update: update,
		admin:  admin,
		stage:  stage,
	}
}

func (e *TGEvent) Command() int {
	if e.SentFrom() != nil && e.SentFrom().UserName == e.admin && e.stage == step.Initialization {
		return command.Initialize
	}

	return command.Unknown
}

func (e *TGEvent) User() int64 {
	if e.FromChat() != nil {
		return e.FromChat().ChatConfig().ChatID
	}

	return 0
}
