package state_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/protocol/command"
	"github.com/wdipax/match/protocol/response"
	"github.com/wdipax/match/state"
)

func TestState(t *testing.T) {
	t.Parallel()

	const (
		admin int64 = iota
	)

	s := state.New()

	var e fakeEvent
	e.command = command.Initialize
	e.user = admin

	r := s.Process(e)
	require.NotEmpty(t, r)

	if assert.Len(t, r.Messages, 2) {
		assert.Equal(t, admin, r.Messages[0].To)
		assert.Equal(t, response.BoysToken, r.Messages[0].Type)

		assert.Equal(t, admin, r.Messages[1].To)
		assert.Equal(t, response.GirlsToken, r.Messages[1].Type)
	}
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
