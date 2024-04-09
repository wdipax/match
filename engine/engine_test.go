package engine_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/engine"
)

func TestEngine(t *testing.T) {
	t.Parallel()

	t.Run("it responds with help message", func(t *testing.T) {
		t.Parallel()

		var tg fakeTelegramHandler

		st := fakeStateHandler{
			helpMsg: "test",
		}

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.Help,
		}

		e.Process(&evt)

		assert.Equal(t, "test", tg.sentMsg)
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

type fakeTelegramHandler struct {
	sentMsg string
}

func (tg *fakeTelegramHandler) Send(msg string) {
	tg.sentMsg = msg
}

type fakeEvent struct {
	action engine.Action
}

func (e *fakeEvent) Command() engine.Action {
	return e.action
}

func (e *fakeEvent) UserID() string {
	return ""
}

type fakeStateHandler struct {
	helpMsg string
}

func (h *fakeStateHandler) Help(userID string) string {
	return h.helpMsg
}

func (h *fakeStateHandler) NewSession(userID string) string {
	return ""
}

func (h *fakeStateHandler) StartMaleRegistration(userID string) string {
	return ""
}
