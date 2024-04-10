// adapter is a bridge between the outer telegram hendler and the inner state of the application.
package adapter

import "github.com/wdipax/match/state"

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

// TODO: this type belongs to the Event implementation, move it there.
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

func (a *Adapter) Process(evt Event) {
	userID := evt.UserID()
	payload := evt.Payload()

	var responses []*state.Response

	switch evt.Command() {
	case Help:
		responses = a.state.Help(userID)
	case NewSession:
		responses = a.state.NewSession(userID)
	case StartMaleRegistration:
		responses = a.state.StartMaleRegistration(userID)
	case EndMaleRegistration:
		responses = a.state.EndMaleRegistration(userID)
	case StartFemaleRegistration:
		responses = a.state.StartFemaleRegistration(userID)
	case EndFemaleRegistration:
		responses = a.state.EndFemaleRegistration(userID)
	case AddTeamMember:
		responses = a.state.AddTeamMember(userID, payload)
	case TeamMemberName:
		responses = a.state.TeamMemberName(userID, payload)
	case TeamMemberNumber:
		responses = a.state.TeamMemberNumber(userID, payload)
	case StartVoting:
		responses = a.state.StartVoting(userID)
	case Vote:
		responses = a.state.Vote(userID, payload)
	case EndSession:
		responses = a.state.EndSession(userID)
	}

	for _, re := range responses {
		a.messenger.Send(re.UserID, re.MSG)
	}
}
