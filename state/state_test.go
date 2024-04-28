package state_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/protocol/command"
	"github.com/wdipax/match/protocol/response"
	"github.com/wdipax/match/state"
	"github.com/wdipax/match/state/view"
)

const (
	admin int64 = iota
	boy1
	boy2
	girl1
	girl2
)

func TestState(t *testing.T) {
	t.Parallel()

	s := &stateHelper{
		State: state.New(),
	}

	initialization(t, s)

	registration(t, s)

	knowingEachOther(t, s)

	voting(t, s)
}

func voting(tb testing.TB, s *stateHelper) {
	tb.Helper()

	vote(tb, s, boy1, number(girl2))
	vote(tb, s, girl2, number(boy1), number(boy2))

	vote(tb, s, boy2, number(girl1))
	revote(tb, s, boy2)
	vote(tb, s, boy2, number(girl2), number(girl1))

	vote(tb, s, girl1, number(boy2))
	revote(tb, s, girl1)

	votingStat(tb, s)

	end(tb, s)
}

func revote(tb testing.TB, s *stateHelper, user int64) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Repeat
	e.user = user

	r := s.Process(e)

	require.Len(tb, r.Messages, 1)
	require.Equal(tb, user, r.Messages[0].To)
	require.NotEmpty(tb, r.Messages[0].Data)
	require.Equal(tb, response.Poll, r.Messages[0].Type)
}

func end(tb testing.TB, s *stateHelper) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Next
	e.user = admin

	r := s.Process(e)

	require.Len(tb, r.Messages, s.cntUsers)

	am := messageFor(tb, r, admin)
	require.Equal(tb, response.End, am.Type)
	require.Empty(tb, am.Data)

	b1 := messageFor(tb, r, boy1)
	require.Equal(tb, response.End, b1.Type)
	require.Contains(tb, b1.Data, contact(girl2))

	b2 := messageFor(tb, r, boy2)
	require.Contains(tb, b2.Data, contact(girl2))

	g1 := messageFor(tb, r, girl1)
	require.Empty(tb, g1.Data)

	g2 := messageFor(tb, r, girl2)
	require.Contains(tb, g2.Data, contact(boy1))
	require.Contains(tb, g2.Data, contact(boy2))
}

func messageFor(tb testing.TB, r *response.Response, user int64) response.Message {
	tb.Helper()

	var m *response.Message

	for _, rm := range r.Messages {
		if rm.To == user {
			m = &rm
			break
		}
	}

	require.NotEmpty(tb, m)

	return *m
}

func votingStat(tb testing.TB, s *stateHelper) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Stat
	e.user = admin

	r := s.Process(e)

	require.Len(tb, r.Messages, 1)
	require.Equal(tb, admin, r.Messages[0].To)
	require.Equal(tb, response.Stat, r.Messages[0].Type)
	require.Contains(tb, "3/4", r.Messages[0].Data)
}

func vote(tb testing.TB, s *stateHelper, user int64, chosen ...int) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Vote
	e.user = user

	var pr view.PollResult

	e.data = pr.Encode(chosen)

	r := s.Process(e)

	require.Len(tb, r.Messages, 1)
	require.Equal(tb, user, r.Messages[0].To)
	require.Equal(tb, response.Success, r.Messages[0].Type)
}

func knowingEachOther(tb testing.TB, s *stateHelper) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Next
	e.user = admin

	r := s.Process(e)

	require.Len(tb, r.Messages, s.cntUsers)

	checker, asserter := messagesFor(tb, admin, boy1, boy2, girl1, girl2)
	checkMessages(r.Messages, checker)
	asserter(tb)
}

func registration(tb testing.TB, s *stateHelper) {
	tb.Helper()

	adminCanNotJoin(tb, s, s.boys)
	adminCanNotJoin(tb, s, s.girls)

	joinGroup(tb, s, boy1, s.boys)
	joinGroup(tb, s, boy2, s.boys)
	joinGroup(tb, s, girl1, s.girls)
	joinGroup(tb, s, girl2, s.girls)

	setName(tb, s, boy1)
	setName(tb, s, boy2)
	setName(tb, s, girl1)
	setName(tb, s, girl2)

	groupStat(tb, s)

	var e fakeEvent
	e.command = command.Next
	e.user = admin

	r := s.Process(e)
	require.NotEmpty(tb, r)
	require.Len(tb, r.Messages, s.cntUsers)

	checker, asserter := messagesFor(tb, admin, boy1, boy2, girl1, girl2)
	checkMessages(r.Messages, checker)
	asserter(tb)
}

func messagesFor(tb testing.TB, users ...int64) (func(response.Message), func(testing.TB)) {
	tb.Helper()

	set := make(map[int64]bool)

	return func(m response.Message) {
			set[m.To] = true
		}, func(tb testing.TB) {
			tb.Helper()

			ok := true

			for _, u := range users {
				if !set[u] {
					tb.Error("no message for", name(u))

					ok = false
				}
			}

			require.True(tb, ok)
		}
}

func checkMessages(msgs []response.Message, checks ...func(response.Message)) {
	for _, m := range msgs {
		for _, c := range checks {
			c(m)
		}
	}
}

func groupStat(tb testing.TB, s *stateHelper) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Stat
	e.user = admin

	r := s.Process(e)
	require.NotEmpty(tb, r)
	require.Len(tb, r.Messages, 2)

	require.Equal(tb, admin, r.Messages[0].To)
	require.Equal(tb, response.ViewBoys, r.Messages[0].Type)
	require.Contains(tb, r.Messages[0].Data, name(boy1))
	require.Contains(tb, r.Messages[0].Data, name(boy2))

	require.Equal(tb, admin, r.Messages[1].To)
	require.Equal(tb, response.ViewGirls, r.Messages[1].Type)
	require.Contains(tb, r.Messages[1].Data, name(girl1))
	require.Contains(tb, r.Messages[1].Data, name(girl2))
}

func setName(tb testing.TB, s *stateHelper, user int64) {
	tb.Helper()

	n := name(user)

	var e fakeEvent
	e.command = command.SetName
	e.user = user
	e.data = n

	r := s.Process(e)
	require.NotEmpty(tb, r)
	require.Len(tb, r.Messages, 1)

	require.Equal(tb, user, r.Messages[0].To)
	require.Contains(tb, r.Messages[0].Data, n)
	require.Equal(tb, response.Success, r.Messages[0].Type)
}

func adminCanNotJoin(tb testing.TB, s *stateHelper, group string) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Join
	e.user = admin
	e.data = group

	r := s.Process(e)
	require.NotEmpty(tb, r)
	require.Len(tb, r.Messages, 1)

	require.Equal(tb, admin, r.Messages[0].To)
	require.Equal(tb, response.RestrictedForAdmin, r.Messages[0].Type)
}

func joinGroup(tb testing.TB, s *stateHelper, user int64, group string) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Join
	e.user = user
	e.data = group
	e.contact = contact(user)

	r := s.Process(e)
	require.NotEmpty(tb, r)
	require.Len(tb, r.Messages, 1)

	require.Equal(tb, user, r.Messages[0].To)

	_, err := strconv.Atoi(r.Messages[0].Data)
	require.NoError(tb, err)

	require.Equal(tb, response.Joined, r.Messages[0].Type)

	s.cntUsers++
}

func initialization(tb testing.TB, s *stateHelper) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Initialize
	e.user = admin

	r := s.Process(e)
	require.NotEmpty(tb, r)
	require.Len(tb, r.Messages, 3)

	require.Equal(tb, admin, r.Messages[0].To)
	require.Equal(tb, response.Control, r.Messages[0].Type)

	require.Equal(tb, admin, r.Messages[1].To)
	require.NotEmpty(tb, r.Messages[1].Data)
	s.boys = r.Messages[1].Data
	require.Equal(tb, response.BoysToken, r.Messages[1].Type)

	require.Equal(tb, admin, r.Messages[2].To)
	require.NotEmpty(tb, r.Messages[2].Data)
	s.girls = r.Messages[2].Data
	require.Equal(tb, response.GirlsToken, r.Messages[2].Type)

	s.cntUsers++
}

type stateHelper struct {
	*state.State

	boys     string
	girls    string
	cntUsers int
}

type fakeEvent struct {
	command int
	user    int64
	data    string
	contact string
}

func (e fakeEvent) Command() int {
	return e.command
}

func (e fakeEvent) User() int64 {
	return e.user
}

func (e fakeEvent) Data() string {
	return e.data
}

func (e fakeEvent) Contact() string {
	return e.contact
}

func name(user int64) string {
	switch user {
	case admin:
		return "admin"
	case boy1:
		return "boy1"
	case boy2:
		return "boy2"
	case girl1:
		return "girl1"
	case girl2:
		return "girl2"
	default:
		return "unknown"
	}
}

func contact(user int64) string {
	return name(user)
}

func number(user int64) int {
	switch user {
	case boy1:
		return 1
	case boy2:
		return 2
	case girl1:
		return 1
	case girl2:
		return 2
	default:
		return 0
	}
}
