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

	t.Run("it prints differrent help messages", func(t *testing.T) {
		t.Parallel()

		t.Run("for admins", func(t *testing.T) {
			t.Parallel()

			// Given there is an admin.
			admin := uuid.NewString()

			// When the admin asks for a help.
			var evt fakeEvent
			evt.userID = admin
			evt.action = engine.Help

			// Then they receive a help message for admins.
			var sh spySessionHandler

			st := state.New(&state.Settings{
				Admins:         []string{admin},
				SessionHandler: &sh,
			})

			e := engine.New(st)

			e.Process(&evt)

			assert.True(t, sh.adminHelpPrinted)
		})

		t.Run("for regular users", func(t *testing.T) {
			t.Parallel()

			// Given there is a user.
			user := uuid.NewString()

			// When the user asks for a help.
			var evt fakeEvent
			evt.userID = user
			evt.action = engine.Help

			// Then they receive a help message for users.
			var sh spySessionHandler

			st := state.New(&state.Settings{
				Admins:         []string{},
				SessionHandler: &sh,
			})

			e := engine.New(st)

			e.Process(&evt)

			assert.True(t, sh.userHelpPrinted)
		})
	})

	t.Run("it starts a new session when requested by an admin", func(t *testing.T) {
		t.Parallel()

		// Given there is an admin.
		admin := uuid.NewString()

		// When the admin starts a new session.
		var evt fakeEvent
		evt.userID = admin
		evt.action = engine.NewSession

		// Then the new session is created.
		var sh spySessionHandler

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
		evt.userID = user
		evt.action = engine.NewSession

		// Then the new session is not created.
		var sh spySessionHandler

		st := state.New(&state.Settings{
			Admins:         []string{},
			SessionHandler: &sh,
		})

		e := engine.New(st)

		e.Process(&evt)

		assert.False(t, sh.sessionCreated)
	})

	t.Run("it does not start a new session when the session is in progress", func(t *testing.T) {
		t.Parallel()

		// Given there is an admin.
		admin := uuid.NewString()

		// When the admin starts a new session twice.
		var evt fakeEvent
		evt.userID = admin
		evt.action = engine.NewSession

		var sh spySessionHandler

		st := state.New(&state.Settings{
			Admins:         []string{admin},
			SessionHandler: &sh,
		})

		e := engine.New(st)

		e.Process(&evt)

		sh.sessionCreated = false

		e.Process(&evt)

		// Then the second time the new session is not created.
		assert.False(t, sh.sessionCreated)
	})
}

type fakeEvent struct {
	userID string
	action engine.Action
}

func (e *fakeEvent) Command() engine.Action {
	return e.action
}

func (e *fakeEvent) UserID() string {
	return e.userID
}

type spySessionHandler struct {
	sessionCreated   bool
	adminHelpPrinted bool
	userHelpPrinted  bool
}

func (h *spySessionHandler) New() string {
	h.sessionCreated = true

	return ""
}

func (spySessionHandler) Delete(id string) {}

func (h *spySessionHandler) Help(admin bool) string {
	if admin {
		h.adminHelpPrinted = true
	} else {
		h.userHelpPrinted = true
	}

	return ""
}
