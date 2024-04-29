package tgresponse

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/adapter/tgcontrol"
	"github.com/wdipax/match/protocol/response"
	"github.com/wdipax/match/protocol/step"
	"github.com/wdipax/match/state/view"
)

func From(r *response.Response, bot string, stage int, admin int64) []tgbotapi.Chattable {
	var res []tgbotapi.Chattable

	for _, m := range r.Messages {
		switch m.Type {
		case response.Control:
			switch stage {
			case step.Registration:
				msg := tgbotapi.NewMessage(m.To, sendLinksTpGuests)

				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton(tgcontrol.Stat(stage)),
					tgbotapi.NewKeyboardButton(tgcontrol.Next(stage)),
				))

				res = append(res, msg)
			}
		case response.BoysToken:
			res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf(linkForBoysTemplate, bot, m.Data)))
		case response.GirlsToken:
			res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf(linkForGirlsTemplate, bot, m.Data)))
		case response.RestrictedForAdmin:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, adminCanNotJoinGroup))
			}
		case response.Restricted:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, canNotJoinGroup))
			}
		case response.Joined:
			res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf(whatIsYourNameTemplate, m.Data)))
		case response.Failed:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, canNotUpdateName))
			}
		case response.Success:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf(nameUpdatedTemplate, m.Data)))
			case step.Voting:
				msg := tgbotapi.NewMessage(m.To, voteReceived)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton(tgcontrol.Repeat(stage)),
				))

				res = append(res, msg)
			}
		case response.ViewBoys:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf(boysGroupInfoTemplate, group(m.Data))))
			}
		case response.ViewGirls:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf(girlsGroupInfoTemplate, group(m.Data))))
			}
		case response.KnowEachother:
			msg := tgbotapi.NewMessage(m.To, startDating)

			if m.To == admin {
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton(tgcontrol.Back(stage)),
					tgbotapi.NewKeyboardButton(tgcontrol.Next(stage)),
				))
			}

			res = append(res, msg)
		case response.BackToRegistration:
			msg := tgbotapi.NewMessage(m.To, backToRegistration)

			if m.To == admin {
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton(tgcontrol.Stat(stage)),
					tgbotapi.NewKeyboardButton(tgcontrol.Next(stage)),
				))
			}

			res = append(res, msg)
		case response.Poll:
			if m.To == admin {
				msg := tgbotapi.NewMessage(admin, votingStarted)

				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton(tgcontrol.Stat(stage)),
					tgbotapi.NewKeyboardButton(tgcontrol.Next(stage)),
				))

				res = append(res, msg)
			} else {
				var p view.Poll

				p.Decode(m.Data)

				poll := tgbotapi.NewPoll(m.To, whoDoYouLike, p.Options...)

				poll.AllowsMultipleAnswers = true

				res = append(res, poll)
			}
		case response.Stat:
			switch stage {
			case step.Voting:
				res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf(votingStatisticsTemplate, m.Data)))
			}
		case response.End:
			switch stage {
			case step.End:
				if m.To == admin {
					msg := tgbotapi.NewMessage(admin, votingEnded)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

					res = append(res, msg)
				} else {
					msg := tgbotapi.NewMessage(m.To, matches(m.Data))
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

					res = append(res, msg)
				}
			}
		}
	}

	return res
}

func group(v string) string {
	if v == "" {
		return noMembers
	}

	return v
}

func matches(v string) string {
	if v == "" {
		return noMatches
	}

	return v
}
