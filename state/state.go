package state

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

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
				Text:   "team registration started",
				Type:   response.TeamRegistration,
			},
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
						Text:   fmt.Sprintf("%d is your number now", i),
					},
				},
			}
		} else {
			err = s.state.session.SetUserName(e.ChatID, e.Text)
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
						Text:   fmt.Sprintf("Hello %q! What is your number?", e.Text),
					},
				},
			}
		}
	}

	if e.FromAdmin && e.Command == event.Statistics {
		stat := func(users []*session.User) string {
			if len(users) == 0 {
				return "is empty righ now"
			}

			slices.SortFunc(users, func(a, b *session.User) int {
				if n := cmp.Compare(a.Number, b.Number); n != 0 {
					return n
				}

				return cmp.Compare(a.Name, b.Name)
			})

			var rows []string

			for _, u := range users {
				rows = append(rows, fmt.Sprintf("%d %s", u.Number, u.Name))
			}

			return strings.Join(rows, "\n")
		}

		return &response.Response{
			Messages: []*response.Message{
				{
					ChatID: e.ChatID,
					Text:   fmt.Sprintf("boys team:\n%s", stat(s.state.session.GetBoys())),
				},
				{
					ChatID: e.ChatID,
					Text:   fmt.Sprintf("girls team:\n%s", stat(s.state.session.GetGirls())),
				},
			},
		}
	}

	if e.FromAdmin && e.Command == event.NextStage {
		const endRegistration = "team registration ended"

		res := &response.Response{
			Messages: []*response.Message{
				{
					ChatID: e.ChatID,
					Text:   endRegistration,
				},
			},
		}

		for _, u := range s.state.session.GetAllUsers() {
			res.Messages = append(res.Messages, &response.Message{
				ChatID: u.ID,
				Text:   endRegistration,
			})
		}

		stg := knowEachOther(s)

		s.state.change(stg)

		return res
	}

	return nil
}

type knowEachOther struct {
	state *State
}

func (s knowEachOther) Process(e *event.Event) *response.Response {
	return nil
}
