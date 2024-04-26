package state_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/protocol/command"
	"github.com/wdipax/match/protocol/response"
	"github.com/wdipax/match/state"
)

const (
	admin int64 = iota
)

func TestState(t *testing.T) {
	t.Parallel()

	s := state.New()

	initialize(t, s)
}

func initialize(tb testing.TB, s *state.State) {
	tb.Helper()

	var e fakeEvent
	e.command = command.Initialize
	e.user = admin

	r := s.Process(e)
	require.NotEmpty(tb, r)
	require.Len(tb, r.Messages, 2)
	require.Equal(tb, admin, r.Messages[0].To)
	require.Equal(tb, response.BoysToken, r.Messages[0].Type)
	require.Equal(tb, admin, r.Messages[1].To)
	require.Equal(tb, response.GirlsToken, r.Messages[1].Type)

}

type fakeEvent struct {
	command int
	user    int64
}

func (e fakeEvent) Command() int {
	return e.command
}

func (e fakeEvent) User() int64 {
	return e.user
}
