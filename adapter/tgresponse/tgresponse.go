package tgresponse

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/adapter/tgcontrol"
	"github.com/wdipax/match/protocol/response"
	"github.com/wdipax/match/protocol/step"
)

func From(r *response.Response, bot string, stage int, admin int64) []tgbotapi.Chattable {
	var res []tgbotapi.Chattable

	for _, m := range r.Messages {
		switch m.Type {
		case response.Control:
			switch stage {
			case step.Registration:
				msg := tgbotapi.NewMessage(m.To, "Отправьте ссылки гостям.")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton(tgcontrol.Stat(stage)),
					tgbotapi.NewKeyboardButton(tgcontrol.Next(stage)),
				))

				res = append(res, msg)
			}
		case response.BoysToken:
			res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf("Ссылка для джентельменов https://t.me/%s?start=%s", bot, m.Data)))
		case response.GirlsToken:
			res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf("Ссылка для леди https://t.me/%s?start=%s", bot, m.Data)))
		case response.RestrictedForAdmin:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, "Вы администратор и по этому не можете присоединиться к группе."))
			}
		case response.Restricted:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, "Не получилось присоединить вас к группе, возможно ссылка не дейсвительна."))
			}
		case response.Joined:
			res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf("Ваш номер: %s\nКак вас зовут?", m.Data)))
		case response.Failed:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, "Не получилось обновить имя, возможно ссылка не дейсвительна."))
			}
		case response.Success:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf("Ваше имя: %s\nЕсли имя не верно вы можете написать его ещё раз.", m.Data)))
			}
		case response.ViewBoys:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf("Джентельмены:\n%s", group(m.Data))))
			}
		case response.ViewGirls:
			switch stage {
			case step.Registration:
				res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf("Леди:\n%s", group(m.Data))))
			}
		case response.KnowEachother:
			msg := tgbotapi.NewMessage(m.To, "Давайте знакомиться.")
			
			if m.To == admin {
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton(tgcontrol.Back(stage)),
					tgbotapi.NewKeyboardButton(tgcontrol.Next(stage)),
				))
			}

			res = append(res, msg)
		}
	}

	return res
}

func group(v string) string {
	if v == "" {
		return "Пока никого."
	}

	return v
}
