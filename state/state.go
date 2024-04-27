package state

import (
	"strconv"

	"github.com/wdipax/match/protocol/command"
	"github.com/wdipax/match/protocol/response"
	"github.com/wdipax/match/protocol/step"
	"github.com/wdipax/match/state/group"
	"github.com/wdipax/match/state/group/member"
	"github.com/wdipax/match/state/view"
)

type Event interface {
	Command() int
	User() int64
	Data() string
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

	stage := initialization{
		State: &state,
	}

	state.stage = stage

	return &state
}

func (s *State) Admin() int64 {
	return s.admin
}

func (s *State) allMembers() []*member.Member {
	return append(s.boys.Members(), s.girls.Members()...)
}

type initialization struct {
	*State
}

func (s initialization) Step() int {
	return step.Initialization
}

func (s initialization) Process(e Event) *response.Response {
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
				Type: response.Control,
			},
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
	if e.Command() == command.Join {
		if e.User() == s.admin {
			return &response.Response{
				Messages: []response.Message{
					{
						To:   s.admin,
						Type: response.RestrictedForAdmin,
					},
				},
			}
		}

		number := s.register(e.Data(), e.User())
		if number == 0 {
			return &response.Response{
				Messages: []response.Message{
					{
						To:   e.User(),
						Type: response.Restricted,
					},
				},
			}
		}

		return &response.Response{
			Messages: []response.Message{
				{
					To:   e.User(),
					Data: strconv.Itoa(number),
					Type: response.Joined,
				},
			},
		}
	}

	if e.Command() == command.SetName {
		ok := s.setName(e.User(), e.Data())
		if !ok {
			return &response.Response{
				Messages: []response.Message{
					{
						To:   e.User(),
						Type: response.Failed,
					},
				},
			}
		}

		return &response.Response{
			Messages: []response.Message{
				{
					To:   e.User(),
					Data: e.Data(),
					Type: response.Success,
				},
			},
		}
	}

	if e.Command() == command.Stat {
		return &response.Response{
			Messages: []response.Message{
				{
					To:   s.admin,
					Data: view.GroupMembers(s.boys.Members(), s.Step()),
					Type: response.ViewBoys,
				},
				{
					To:   s.admin,
					Data: view.GroupMembers(s.girls.Members(), s.Step()),
					Type: response.ViewGirls,
				},
			},
		}
	}

	if e.Command() == command.Next {
		var msgs []response.Message

		msgs = append(msgs, response.Message{
			To:   s.admin,
			Type: response.KnowEachother,
		})


		for _, m := range s.allMembers() {
			msgs = append(msgs, response.Message{
				To:   m.User,
				Type: response.KnowEachother,
			})
		}

		s.stage = knowEachother(s)

		return &response.Response{
			Messages: msgs,
		}
	}

	return nil
}

func (s registration) register(token string, user int64) int {
	switch {
	case token == s.boys.ID:
		return s.boys.Add(user)
	case token == s.girls.ID:
		return s.girls.Add(user)
	default:
		return 0
	}
}

func (s registration) setName(user int64, name string) bool {
	switch {
	case s.boys.HasMember(user):
		s.boys.SetName(user, name)

		return true
	case s.girls.HasMember(user):
		s.girls.SetName(user, name)

		return true
	default:
		return false
	}
}

type knowEachother struct {
	*State
}

func (s knowEachother) Step() int {
	return step.KnowEachother
}

func (s knowEachother) Process(e Event) *response.Response {
	return nil
}
