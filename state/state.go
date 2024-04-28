package state

import (
	"fmt"
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
	Contact() string
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

func (s *State) allUsers() []int64 {
	var users []int64

	users = append(users, s.admin)
	users = append(users, s.boys.Users()...)
	users = append(users, s.girls.Users()...)

	return users
}

func (s *State) messagesTo(messageType int, data string, users ...int64) []response.Message {
	var msgs []response.Message

	for _, u := range users {
		msgs = append(msgs, response.Message{
			To:   u,
			Data: data,
			Type: messageType,
		})
	}

	return msgs
}

func (s *State) getMember(user int64) *member.Member {
	m := s.boys.Member(user)
	if m != nil {
		return m
	}

	return s.girls.Member(user)
}

func (s *State) allMembers() []*member.Member {
	return append(s.boys.Members(), s.girls.Members()...)
}

func (s *State) getMatchesFor(user int64) []*member.Member {
	m := s.getMember(user)
	if m == nil {
		return nil
	}

	opposite := s.opposite(user)

	choices := s.choices(opposite, m.Choosen)

	var matches []*member.Member

	for _, c := range choices {
		for _, match := range s.choices(s.opposite(c.User), c.Choosen) {
			if match.User == user {
				matches = append(matches, c)
			}
		}
	}

	return matches
}

func (s *State) choices(g *group.Group, numbers []int) []*member.Member {
	var acc []*member.Member

	for _, number := range numbers {
		acc = append(acc, g.MemberBy(number))
	}

	return acc
}

func (s *State) opposite(user int64) *group.Group {
	switch {
	case s.boys.HasMember(user):
		return s.girls
	case s.girls.HasMember(user):
		return s.boys
	default:
		return nil
	}
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

		number := s.register(e.Data(), e.User(), e.Contact())
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
					Data: view.GroupMembers(s.boys.Members()),
					Type: response.ViewBoys,
				},
				{
					To:   s.admin,
					Data: view.GroupMembers(s.girls.Members()),
					Type: response.ViewGirls,
				},
			},
		}
	}

	if e.Command() == command.Next {
		s.stage = knowEachother(s)

		return &response.Response{
			Messages: s.messagesTo(response.KnowEachother, "", s.allUsers()...),
		}
	}

	return nil
}

func (s registration) register(token string, user int64, contact string) int {
	switch {
	case token == s.boys.ID:
		return s.boys.Add(user, contact)
	case token == s.girls.ID:
		return s.girls.Add(user, contact)
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
	if e.Command() == command.Back {
		s.stage = registration(s)

		return &response.Response{
			Messages: s.messagesTo(response.BackToRegistration, "", s.allUsers()...),
		}
	}

	if e.Command() == command.Next {
		var msgs []response.Message

		msgs = append(msgs, response.Message{
			To:   s.admin,
			Type: response.Poll,
		})

		msgs = append(msgs, s.messagesTo(response.Poll, view.FormPoll(s.girls.Members()), s.boys.Users()...)...)
		msgs = append(msgs, s.messagesTo(response.Poll, view.FormPoll(s.boys.Members()), s.girls.Users()...)...)

		s.stage = voting(s)

		return &response.Response{
			Messages: msgs,
		}
	}

	return nil
}

type voting struct {
	*State
}

func (s voting) Step() int {
	return step.Voting
}

func (s voting) Process(e Event) *response.Response {
	if e.Command() == command.Vote {
		var r view.PollResult

		r.Decode(e.Data())

		m := s.getMember(e.User())

		m.Choosen = r.Choosen

		return &response.Response{
			Messages: []response.Message{
				{
					To:   e.User(),
					Type: response.Success,
				},
			},
		}
	}

	if e.Command() == command.Repeat {
		m := s.getMember(e.User())
		m.Choosen = nil
		
		opposite := s.opposite(e.User())
		
		return &response.Response{
			Messages: []response.Message{
				{
					To: e.User(),
					Data: view.FormPoll(opposite.Members()),
					Type: response.Poll,
				},
			},
		}
	}

	if e.Command() == command.Stat {
		am := s.allMembers()

		var voted int

		for _, m := range am {
			if len(m.Choosen) > 0 {
				voted++
			}
		}

		return &response.Response{
			Messages: []response.Message{
				{
					To:   s.admin,
					Data: fmt.Sprintf("%d/%d", voted, len(am)),
					Type: response.Stat,
				},
			},
		}
	}

	if e.Command() == command.Next {
		var msgs []response.Message

		for _, user := range s.allUsers() {
			matches := s.getMatchesFor(user)

			msgs = append(msgs, response.Message{
				To:   user,
				Data: view.Matches(matches),
				Type: response.End,
			})
		}

		s.stage = end(s)

		return &response.Response{
			Messages: msgs,
		}
	}

	return nil
}

type end struct {
	*State
}

func (s end) Step() int {
	return step.End
}

func (s end) Process(e Event) *response.Response {
	return nil
}
