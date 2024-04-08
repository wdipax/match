package engine_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/engine"
)

func TestEngine(t *testing.T) {
	t.Parallel()

	t.Run("it prints help message", func(t *testing.T) {
		t.Parallel()

		userID := uuid.NewString()

		var evt fakeEvent
		evt.userID = userID
		evt.action = engine.Help

		var sh fakeStateHandler

		e := engine.New(&sh)

		e.Process(&evt)

		assert.True(t, sh.helpCalled)
	})

	// t.Run("it starts a new session when requested by an admin", func(t *testing.T) {
	// 	t.Parallel()

	// 	// Given there is an admin.
	// 	admin := uuid.NewString()

	// 	// When the admin starts a new session.
	// 	var evt fakeEvent
	// 	evt.userID = admin
	// 	evt.action = engine.NewSession

	// 	// Then the new session is created.
	// 	var sh spySessionHandler

	// 	st := state.New(&state.Settings{
	// 		Admins:         []string{admin},
	// 		SessionHandler: &sh,
	// 	})

	// 	e := engine.New(st)

	// 	e.Process(&evt)

	// 	assert.True(t, sh.sessionCreated)
	// })

	// t.Run("it does not start a new session when requested by not an admin", func(t *testing.T) {
	// 	t.Parallel()

	// 	// Given there is a user.
	// 	user := uuid.NewString()

	// 	// When the user starts a new session.
	// 	var evt fakeEvent
	// 	evt.userID = user
	// 	evt.action = engine.NewSession

	// 	// Then the new session is not created.
	// 	var sh spySessionHandler

	// 	st := state.New(&state.Settings{
	// 		Admins:         []string{},
	// 		SessionHandler: &sh,
	// 	})

	// 	e := engine.New(st)

	// 	e.Process(&evt)

	// 	assert.False(t, sh.sessionCreated)
	// })

	// t.Run("it does not start a new session when the session is in progress", func(t *testing.T) {
	// 	t.Parallel()

	// 	// Given there is an admin.
	// 	admin := uuid.NewString()

	// 	// When the admin starts a new session twice.
	// 	var evt fakeEvent
	// 	evt.userID = admin
	// 	evt.action = engine.NewSession

	// 	var sh spySessionHandler

	// 	st := state.New(&state.Settings{
	// 		Admins:         []string{admin},
	// 		SessionHandler: &sh,
	// 	})

	// 	e := engine.New(st)

	// 	e.Process(&evt)

	// 	sh.sessionCreated = false

	// 	e.Process(&evt)

	// 	// Then the second time the new session is not created.
	// 	assert.False(t, sh.sessionCreated)
	// })

	// t.Run("it starts male registration", func(t *testing.T) {
	// 	t.Parallel()

	// 	// Given there is an admin.
	// 	admin := uuid.NewString()

	// 	// When the admin starts a new session.
	// 	var evt fakeEvent
	// 	evt.userID = admin
	// 	evt.action = engine.NewSession

	// 	var sh spySessionHandler

	// 	st := state.New(&state.Settings{
	// 		Admins:         []string{admin},
	// 		SessionHandler: &sh,
	// 	})

	// 	e := engine.New(st)

	// 	e.Process(&evt)

	// 	// And starts a male team registration.
	// 	evt.action = engine.StartMaleRegistration

	// 	e.Process(&evt)

	// 	// Then the male registration is started.
	// 	assert.True(t, sh.maleRegistrationStarted)
	// })
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

type fakeStateHandler struct {
	helpCalled bool
}

func (h *fakeStateHandler) Help(userID string) string {
	h.helpCalled = true

	return ""
}

func (h *fakeStateHandler) NewSession(userID string) error {
	return nil
}

func (h *fakeStateHandler) StartMaleRegistration(userID string) error {
	return nil
}

type spySessionHandler struct {
	sessionCreated          bool
	adminHelpPrinted        bool
	userHelpPrinted         bool
	maleRegistrationStarted bool
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

func (h *spySessionHandler) StartRegistration(teamID string) error {
	h.maleRegistrationStarted = true

	return nil
}
