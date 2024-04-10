// engine is a bridge between the outer telegram hendler and the inner logic of the application.
package engine

import "github.com/wdipax/match/state"

type Engine struct {
	messenger MessengerHandler
	state     StateHandler
}

type MessengerHandler interface {
	Send(userID string, msg string)
}

type StateHandler interface {
	Help(userID string) []*state.Response
	NewSession(userID string) []*state.Response
	StartMaleRegistration(userID string) []*state.Response
	EndMaleRegistration(userID string) []*state.Response
	StartFemaleRegistration(userID string) []*state.Response
	EndFemaleRegistration(userID string) []*state.Response
	AddTeamMember(userID string, teamID string) []*state.Response
	TeamMemberName(userID string, name string) []*state.Response
	TeamMemberNumber(userID string, number string) []*state.Response
	StartVoting(userID string) []*state.Response
	Vote(userID string, poll string) []*state.Response
	EndSession(userID string) []*state.Response
}

func New(telegram MessengerHandler, state StateHandler) *Engine {
	return &Engine{
		messenger: telegram,
		state:     state,
	}
}

// Action represenst an event action.
type Action int

const (
	Unknown Action = iota
	Help
	NewSession
	EndSession
	StartMaleRegistration
	EndMaleRegistration
	StartFemaleRegistration
	EndFemaleRegistration
	AddTeamMember
	TeamMemberName
	TeamMemberNumber
	StartVoting
	Vote
)

// Event represents an event from the telegram.
type Event interface {
	Command() Action
	UserID() string
	Payload() string
}

func (e *Engine) Process(evt Event) {
	userID := evt.UserID()
	payload := evt.Payload()

	var responses []*state.Response

	switch evt.Command() {
	case Help:
		responses = e.state.Help(userID)
	case NewSession:
		responses = e.state.NewSession(userID)
	case StartMaleRegistration:
		responses = e.state.StartMaleRegistration(userID)
	case EndMaleRegistration:
		responses = e.state.EndMaleRegistration(userID)
	case StartFemaleRegistration:
		responses = e.state.StartFemaleRegistration(userID)
	case EndFemaleRegistration:
		responses = e.state.EndFemaleRegistration(userID)
	case AddTeamMember:
		responses = e.state.AddTeamMember(userID, payload)
	case TeamMemberName:
		responses = e.state.TeamMemberName(userID, payload)
	case TeamMemberNumber:
		responses = e.state.TeamMemberNumber(userID, payload)
	case StartVoting:
		responses = e.state.StartVoting(userID)
	case Vote:
		responses = e.state.Vote(userID, payload)
	case EndSession:
		responses = e.state.EndSession(userID)
	}

	for _, re := range responses {
		e.messenger.Send(re.UserID, re.MSG)
	}
}
