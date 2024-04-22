package state

import (
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/response"
	"github.com/wdipax/match/session"
)

type stage interface {
	Process(e *event.Event) *response.Response
}

type State struct {
	stage

	session *session.Session
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
	if !e.FromAdmin {
		return nil
	}

	ss := session.New()

	s.state.session = ss

	boysID := ss.CreateBoysTeam()

	girlsID := ss.CreateGirlsTeam()

	stg := teamsRegistration(s)

	s.state.change(stg)

	return &response.Response{
		Messages: []*response.Message{
			{
				Text: boysID,
			},
			{
				Text: girlsID,
			},
		},
	}
}

type teamsRegistration struct {
	state *State
}

func (s teamsRegistration) Process(e *event.Event) *response.Response {
	if e.EndTeamRegistration {
		stg := knowEachOther(s)

		s.state.change(stg)

		return nil
	}

	return nil
}

type knowEachOther struct {
	state *State
}

func (s knowEachOther) Process(e *event.Event) *response.Response {
	return nil
}
