package state

import (
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/response"
)

type stage interface {
	Process(e *event.Event) *response.Response
}

type State struct {
	stage
}

func New() *State {
	var s State

	st := waitForAdmin{
		state: &s,
	}

	s.stage = st

	return &s
}

func (s *State) change(st stage) {
	s.stage = st
}

type waitForAdmin struct {
	state *State
}

func (s waitForAdmin) Process(e *event.Event) *response.Response {
	if !e.FromAdmin() {
		return nil
	}

	stg := teamsRegistration(s)

	s.state.change(stg)

	// TODO: return links for joining teams.
	return nil
}

type teamsRegistration struct {
	state *State
}

func (s teamsRegistration) Process(e *event.Event) *response.Response {
	return nil
}
