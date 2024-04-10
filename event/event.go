// event converts an incomming notification from messenger to the inner domain object.
package event

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Event struct {
	update *tgbotapi.Update
}

func From(update *tgbotapi.Update) *Event {
	return &Event{
		update: update,
	}
}

type Type int

const (
	Unknown Type = iota
	Input
	Help
	NewSession
	StartMaleRegistration
	EndMaleRegistration
	StartFemaleRegistration
	EndFemaleRegistration
	ChangeUserName
	ChangeUserNumber
	StartVoting
	EndSession
)

func (e *Event) Command() Type {
	msg := e.update.Message
	if msg == nil {
		return Unknown
	}

	if !strings.HasPrefix(msg.Text, "/") {
		return Input
	}

	switch msg.Text {
	case "/help":
		return Help
	case "/new_session":
		return NewSession
	case "/start_male_registration":
		return StartMaleRegistration
	case "/end_male_registration":
		return EndMaleRegistration
	case "/start_female_registration":
		return StartFemaleRegistration
	case "/end_female_registration":
		return EndFemaleRegistration
	case "/change_user_name":
		return ChangeUserName
	case "/change_user_number":
		return ChangeUserNumber
	case "/start_voting":
		return StartVoting
	case "/end_session":
		return EndSession
	default:
		return Unknown
	}
}

func (e *Event) UserID() string {
	msg := e.update.Message
	if msg == nil {
		return ""
	}

	user := msg.From
	if user == nil {
		return ""
	}

	return strconv.Itoa(int(user.ID))
}

func (e *Event) Payload() string {
	msg := e.update.Message
	if msg == nil {
		return ""
	}

	return msg.Text
}
