package state

import (
	"fmt"
	"strconv"

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
				ChatID: e.ChatID,
				Text:   boysID,
				Type:   response.BoysLink,
			},
			{
				ChatID: e.ChatID,
				Text:   girlsID,
				Type:   response.GirlsLink,
			},
		},
	}
}

type teamsRegistration struct {
	state *State
}

func (s teamsRegistration) Process(e *event.Event) *response.Response {
	if e.FromAdmin && e.TeamID != "" {
		return &response.Response{
			Messages: []*response.Message{
				{
					ChatID: e.ChatID,
					Text:   "you can not join a team, you are an admin",
				},
			},
		}
	}

	if !e.FromAdmin && e.TeamID != "" {
		ok := s.state.session.JoinTeam(e.TeamID, e.ChatID)
		if !ok {
			return &response.Response{
				Messages: []*response.Message{
					{
						ChatID: e.ChatID,
						Text:   "no such team",
					},
				},
			}
		}

		return &response.Response{
			Messages: []*response.Message{
				{
					ChatID: e.ChatID,
					Text:   "what is your name?",
				},
			},
		}
	}

	if !e.FromAdmin && e.TeamID == "" {
		i, err := strconv.Atoi(e.Text)
		if err == nil {
			err = s.state.session.SetUserNumber(e.ChatID, i)
			if err != nil {
				return &response.Response{
					Messages: []*response.Message{
						{
							ChatID: e.ChatID,
							Text:   err.Error(),
						},
					},
				}
			}

			return &response.Response{
				Messages: []*response.Message{
					{
						ChatID: e.ChatID,
						Text: fmt.Sprintf("%d is your number now", i),
					},
				},
			}
		} else {
			// s.state.session.SetUserName(e.ChatID, e.Text)
		}
	}

	// if e.EndTeamRegistration {
	// 	stg := knowEachOther(s)

	// 	s.state.change(stg)

	// 	return nil
	// }

	return nil
}

type knowEachOther struct {
	state *State
}

func (s knowEachOther) Process(e *event.Event) *response.Response {
	return nil
}
