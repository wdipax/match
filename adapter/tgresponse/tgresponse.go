package tgresponse

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/protocol/response"
)

func From(r *response.Response, bot string) []tgbotapi.Chattable {
	var res []tgbotapi.Chattable

	for _, m := range r.Messages {
		switch m.Type {
		case response.BoysToken:
			res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf("Ссылка для джентельменов https://t.me/%s?start=%s", bot, m.Data)))
		case response.GirlsToken:
			res = append(res, tgbotapi.NewMessage(m.To, fmt.Sprintf("Ссылка для леди https://t.me/%s?start=%s", bot, m.Data)))
		}
	}

	return res
}
