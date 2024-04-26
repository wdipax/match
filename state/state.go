package state

import (
	"github.com/wdipax/match/protocol/command"
	"github.com/wdipax/match/protocol/response"
	"github.com/wdipax/match/protocol/step"
	"github.com/wdipax/match/state/group"
)

type Event interface {
	Command() int
	User() int64
}

type stage interface {
	Process(e Event) *response.Response
	Step() int
}

type State struct {
	stage

	admin int64
	boys  *group.Group
	girls *group.Group
}

func New() *State {
	var state State

	stage := initial{
		State: &state,
	}

	state.stage = stage

	return &state
}

type initial struct {
	*State
}

func (s initial) Step() int {
	return step.Initialization
}

func (s initial) Process(e Event) *response.Response {
	if e.Command() != command.Initialize {
		return nil
	}

	s.admin = e.User()

	s.boys = group.New()
	s.girls = group.New()

	s.stage = registration(s)

	return &response.Response{
		Messages: []response.Message{
			{
				To:   s.admin,
				Data: s.boys.ID,
				Type: response.BoysToken,
			},
			{
				To:   s.admin,
				Data: s.girls.ID,
				Type: response.GirlsToken,
			},
		},
	}
}

type registration struct {
	*State
}

func (s registration) Step() int {
	return step.Registration
}

func (s registration) Process(e Event) *response.Response {
	return nil
}
