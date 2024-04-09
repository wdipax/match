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

	t.Run("it starts a new session", func(t *testing.T) {
		t.Parallel()

		var tg fakeTelegramHandler

		st := fakeStateHandler{
			newSessionMsg: "test",
		}

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.NewSession,
		}

		e.Process(&evt)

		assert.Equal(t, "test", tg.sentMsg)
	})

	t.Run("it starts male registration", func(t *testing.T) {
		t.Parallel()

		var tg fakeTelegramHandler

		st := fakeStateHandler{
			startMaleRegistrationMsg: "test",
		}

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.StartMaleRegistration,
		}

		e.Process(&evt)

		assert.Equal(t, "test", tg.sentMsg)
	})

	t.Run("it ends male registration", func(t *testing.T) {
		t.Parallel()

		var tg fakeTelegramHandler

		st := fakeStateHandler{
			endMaleRegistrationMsg: "test",
		}

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.EndMaleRegistration,
		}

		e.Process(&evt)

		assert.Equal(t, "test", tg.sentMsg)
	})

	t.Run("it starts female registration", func(t *testing.T) {
		t.Parallel()

		var tg fakeTelegramHandler

		st := fakeStateHandler{
			startFemaleRegistrationMsg: "test",
		}

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.StartFemaleRegistration,
		}

		e.Process(&evt)

		assert.Equal(t, "test", tg.sentMsg)
	})

	t.Run("it ends female registration", func(t *testing.T) {
		t.Parallel()

		var tg fakeTelegramHandler

		st := fakeStateHandler{
			endFemaleRegistrationMsg: "test",
		}

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.EndFemaleRegistration,
		}

		e.Process(&evt)

		assert.Equal(t, "test", tg.sentMsg)
	})
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
	helpMsg                    string
	newSessionMsg              string
	startMaleRegistrationMsg   string
	endMaleRegistrationMsg     string
	startFemaleRegistrationMsg string
	endFemaleRegistrationMsg   string
}

func (h *fakeStateHandler) Help(userID string) string {
	return h.helpMsg
}

func (h *fakeStateHandler) NewSession(userID string) string {
	return h.newSessionMsg
}

func (h *fakeStateHandler) StartMaleRegistration(userID string) string {
	return h.startMaleRegistrationMsg
}

func (h *fakeStateHandler) EndMaleRegistration(userID string) string {
	return h.endMaleRegistrationMsg
}

func (h *fakeStateHandler) StartFemaleRegistration(userID string) string {
	return h.startFemaleRegistrationMsg
}

func (h *fakeStateHandler) EndFemaleRegistration(userID string) string {
	return h.endFemaleRegistrationMsg
}
