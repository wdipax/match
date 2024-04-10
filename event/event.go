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
	case "/start_male_registration":
		return StartMaleRegistration
	case "/end_male_registration":
		return EndMaleRegistration
	case "/start_female_registration":
		return StartFemaleRegistration
	case "/end_female_registration":
		return EndFemaleRegistration
	case "/add_team_member":
		return AddTeamMember
	case "/team_member_name":
		return TeamMemberName
	case "/team_member_number":
		return TeamMemberNumber
	case "/start_voting":
		return StartVoting
	case "/vote":
		return Vote
	case "/end_session":
		return EndSession
	default:
		return Unknown
	}
}

// func (e *Event) UserID() string {

// }

// func (e *Event) Payload() string {

// }
