// adapter is a bridge between the outer telegram hendler and the inner state of the application.
package adapter

import (
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/state"
)

type Messenger interface {
	Send(userID string, msg string)
}

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

type State interface {
	Input(userID string, payload string) []*state.Response
	Help(userID string) []*state.Response
	NewSession(userID string) []*state.Response
	StartMaleRegistration(userID string) []*state.Response
	EndMaleRegistration(userID string) []*state.Response
	StartFemaleRegistration(userID string) []*state.Response
	EndFemaleRegistration(userID string) []*state.Response
	ChangeUserName(userID string) []*state.Response
	ChangeUserNumber(userID string) []*state.Response
	StartVoting(userID string) []*state.Response
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
	case event.Input:
		responses = a.state.Input(userID, payload)
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
	case event.ChangeUserName:
		responses = a.state.ChangeUserName(userID)
	case event.ChangeUserNumber:
		responses = a.state.ChangeUserNumber(userID)
	case event.StartVoting:
		responses = a.state.StartVoting(userID)
	case event.EndSession:
		responses = a.state.EndSession(userID)
	}

	for _, re := range responses {
		a.messenger.Send(re.UserID, re.MSG)
	}
}
