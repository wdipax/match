package tgevent

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/adapter/tgcontrol"
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
	switch {
	case e.SentFrom() != nil && e.fromAdmin() && e.stage == step.Initialization:
		return command.Initialize
	case e.Message != nil && e.Message.Command() == "start":
		return command.Join
	case e.Message != nil && !e.fromAdmin() && e.stage == step.Registration:
		return command.SetName
	case e.Message != nil && e.fromAdmin() && e.stage != step.Initialization:
		switch e.Message.Text {
		case tgcontrol.Stat(e.stage):
			return command.Stat
		case tgcontrol.Next(e.stage):
			return command.Next
		default:
			return command.Unknown
		}
	default:
		return command.Unknown
	}
}

func (e *TGEvent) fromAdmin() bool {
	return e.Message != nil && e.SentFrom().UserName == e.admin
}

func (e *TGEvent) User() int64 {
	if e.FromChat() != nil {
		return e.FromChat().ChatConfig().ChatID
	}

	return 0
}

func (e *TGEvent) Data() string {
	if e.Message != nil {
		switch {
		case e.Message.IsCommand():
			return e.Message.CommandArguments()
		default:
			return e.Message.Text
		}
	}

	return ""
}
