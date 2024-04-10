// adapter is a bridge between the outer telegram hendler and the inner state of the application.
package adapter

import (
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/state"
)

type Adapter struct {
	messenger Messenger
	state     State
}

func New(messenger Messenger, state State) *Adapter {
	return &Adapter{
		messenger: messenger,
		state:     state,
	}
}

type Messenger interface {
	Send(userID string, msg string)
}

type State interface {
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

type Event interface {
	Command() event.Type
	UserID() string
	Payload() string
}

func (a *Adapter) Process(evt Event) {
	userID := evt.UserID()
	payload := evt.Payload()

	var responses []*state.Response

	switch evt.Command() {
	case event.Help:
		responses = a.state.Help(userID)
	case event.NewSession:
		responses = a.state.NewSession(userID)
	case event.StartMaleRegistration:
		responses = a.state.StartMaleRegistration(userID)
	case event.EndMaleRegistration:
		responses = a.state.EndMaleRegistration(userID)
	case event.StartFemaleRegistration:
		responses = a.state.StartFemaleRegistration(userID)
	case event.EndFemaleRegistration:
		responses = a.state.EndFemaleRegistration(userID)
	case event.AddTeamMember:
		responses = a.state.AddTeamMember(userID, payload)
	case event.TeamMemberName:
		responses = a.state.TeamMemberName(userID, payload)
	case event.TeamMemberNumber:
		responses = a.state.TeamMemberNumber(userID, payload)
	case event.StartVoting:
		responses = a.state.StartVoting(userID)
	case event.Vote:
		responses = a.state.Vote(userID, payload)
	case event.EndSession:
		responses = a.state.EndSession(userID)
	}

	for _, re := range responses {
		a.messenger.Send(re.UserID, re.MSG)
	}
}
