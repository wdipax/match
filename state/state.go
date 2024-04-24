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
			// TODO: why do not set the number at the registration moment?
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

			sortUsers(users)

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
					Type:   response.KnowEachOther,
				},
			},
		}

		res.Messages = append(res.Messages, userMessages(s.state.session.GetAllUsers(), endRegistration)...)

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
	if e.FromAdmin && e.Command == event.PreviousStage {
		const restartRegistration = "team registration reopened"

		res := &response.Response{
			Messages: []*response.Message{
				{
					ChatID: e.ChatID,
					Text:   restartRegistration,
					Type:   response.TeamRegistration,
				},
			},
		}

		res.Messages = append(res.Messages, userMessages(s.state.session.GetAllUsers(), restartRegistration)...)

		stg := teamsRegistration(s)

		s.state.change(stg)

		return res
	}

	if e.FromAdmin && e.Command == event.NextStage {
		const startVoting = "it is time to choose"

		res := &response.Response{
			Messages: []*response.Message{
				{
					ChatID: e.ChatID,
					Text:   startVoting,
					Type:   response.Voting,
				},
			},
		}

		boysList := pollList(s.state.session.GetBoys())

		girlsList := pollList(s.state.session.GetGirls())

		res.Messages = append(res.Messages, userMessages(s.state.session.GetGirls(), poll(startVoting, boysList))...)

		res.Messages = append(res.Messages, userMessages(s.state.session.GetBoys(), poll(startVoting, girlsList))...)

		stg := voting(s)

		s.state.change(stg)

		return res
	}

	return nil
}

type voting struct {
	state *State
}

func (s voting) Process(e *event.Event) *response.Response {

	return nil
}

func poll(header string, options []string) string {
	return strings.Join(append([]string{header}, options...), "\n")
}

func pollList(users []*session.User) []string {
	sortUsers(users)

	var res []string

	for _, u := range users {
		res = append(res, voteInfo(u))
	}

	return res
}

func voteInfo(u *session.User) string {
	return fmt.Sprintf("%d %s", u.Number, u.Name)
}

func sortUsers(users []*session.User) {
	slices.SortFunc(users, func(a, b *session.User) int {
		if n := cmp.Compare(a.Number, b.Number); n != 0 {
			return n
		}

		return cmp.Compare(a.Name, b.Name)
	})
}

func userMessages(users []*session.User, text string) []*response.Message {
	var res []*response.Message

	for _, u := range users {
		res = append(res, &response.Message{
			ChatID: u.ID,
			Text:   text,
		})
	}

	return res
}
