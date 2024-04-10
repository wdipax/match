// event converts an incomming notification from messenger to the inner domain object.
package event

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
	Help
	NewSession
	StartMaleRegistration
	EndMaleRegistration
	StartFemaleRegistration
	EndFemaleRegistration
	AddTeamMember
	TeamMemberName
	TeamMemberNumber
	StartVoting
	Vote
	EndSession
)

func (e *Event) Command() Type {
	msg := e.update.Message
	if msg == nil {
		return Unknown
	}

	switch msg.Text {
	case "/help":
		return Help
	case "/new_session":
		return NewSession
	default:
		return Unknown
	}
}

// func (e *Event) UserID() string {

// }

// func (e *Event) Payload() string {

// }
