package engine_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/engine"
	"github.com/wdipax/match/state"
)

func TestEngine(t *testing.T) {
	t.Parallel()

	t.Run("it starts a new session when requested by an admin", func(t *testing.T) {
		t.Parallel()

		// Given there is an admin.
		admin := uuid.NewString()

		// When the admin starts a new session.
		var evt fakeEvent
		evt.userID = admin

		// Then the new session is created.
		var sh fakeSessionHandler

		st := state.New(&state.Settings{
			Admins:         []string{admin},
			SessionHandler: &sh,
		})

		e := engine.New(st)

		e.Process(&evt)

		assert.True(t, sh.sessionCreated)
	})

	t.Run("it does not start a new session when requested by not an admin", func(t *testing.T) {
		t.Parallel()

		// Given there is a user.
		user := uuid.NewString()

		// When the user starts a new session.
		var evt fakeEvent

		// Then the new session is not created.
		var sh fakeSessionHandler

		st := state.New(&state.Settings{
			Admins:         []string{user},
			SessionHandler: &sh,
		})

		e := engine.New(st)

		e.Process(&evt)

		assert.False(t, sh.sessionCreated)
	})
}

type fakeEvent struct {
	userID string
}

func (fakeEvent) Command() engine.Action {
	return engine.NewSession
}

func (e *fakeEvent) UserID() string {
	return e.userID
}

type fakeSessionHandler struct {
	sessionCreated bool
}

func (f *fakeSessionHandler) New() string {
	f.sessionCreated = true

	return ""
}

func (fakeSessionHandler) Delete(id string) {}
