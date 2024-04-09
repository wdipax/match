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

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.Help,
			userID: "user",
		}

		e.Process(&evt)

		assert.Equal(t, "help user", tg.sentMsg)
	})

	// t.Run("it starts a new session", func(t *testing.T) {
	// 	t.Parallel()

	// 	var tg fakeTelegramHandler

	// 	st := fakeStateHandler{
	// 		newSessionMsg: "test",
	// 	}

	// 	e := engine.New(&tg, &st)

	// 	evt := fakeEvent{
	// 		action: engine.NewSession,
	// 	}

	// 	e.Process(&evt)

	// 	assert.Equal(t, "test", tg.sentMsg)
	// })

	// t.Run("it starts male registration", func(t *testing.T) {
	// 	t.Parallel()

	// 	var tg fakeTelegramHandler

	// 	st := fakeStateHandler{
	// 		startMaleRegistrationMsg: "test",
	// 	}

	// 	e := engine.New(&tg, &st)

	// 	evt := fakeEvent{
	// 		action: engine.StartMaleRegistration,
	// 	}

	// 	e.Process(&evt)

	// 	assert.Equal(t, "test", tg.sentMsg)
	// })

	// t.Run("it ends male registration", func(t *testing.T) {
	// 	t.Parallel()

	// 	var tg fakeTelegramHandler

	// 	st := fakeStateHandler{
	// 		endMaleRegistrationMsg: "test",
	// 	}

	// 	e := engine.New(&tg, &st)

	// 	evt := fakeEvent{
	// 		action: engine.EndMaleRegistration,
	// 	}

	// 	e.Process(&evt)

	// 	assert.Equal(t, "test", tg.sentMsg)
	// })

	// t.Run("it starts female registration", func(t *testing.T) {
	// 	t.Parallel()

	// 	var tg fakeTelegramHandler

	// 	st := fakeStateHandler{
	// 		startFemaleRegistrationMsg: "test",
	// 	}

	// 	e := engine.New(&tg, &st)

	// 	evt := fakeEvent{
	// 		action: engine.StartFemaleRegistration,
	// 	}

	// 	e.Process(&evt)

	// 	assert.Equal(t, "test", tg.sentMsg)
	// })

	// t.Run("it ends female registration", func(t *testing.T) {
	// 	t.Parallel()

	// 	var tg fakeTelegramHandler

	// 	st := fakeStateHandler{
	// 		endFemaleRegistrationMsg: "test",
	// 	}

	// 	e := engine.New(&tg, &st)

	// 	evt := fakeEvent{
	// 		action: engine.EndFemaleRegistration,
	// 	}

	// 	e.Process(&evt)

	// 	assert.Equal(t, "test", tg.sentMsg)
	// })

	// t.Run("it adds team member", func(t *testing.T) {
	// 	t.Parallel()

	// 	var tg fakeTelegramHandler

	// 	st := fakeStateHandler{
	// 		addTeamMemberMsg: "test",
	// 	}

	// 	e := engine.New(&tg, &st)

	// 	evt := fakeEvent{
	// 		action: engine.AddTeamMember,
	// 	}

	// 	e.Process(&evt)

	// 	assert.Equal(t, "test", tg.sentMsg)
	// })
}

type fakeTelegramHandler struct {
	sentMsg string
}

func (tg *fakeTelegramHandler) Send(msg string) {
	tg.sentMsg = msg
}

type fakeEvent struct {
	action  engine.Action
	userID  string
	payload string
}

func (e *fakeEvent) Command() engine.Action {
	return e.action
}

func (e *fakeEvent) UserID() string {
	return e.userID
}

func (e *fakeEvent) Payload() string {
	return e.payload
}

type fakeStateHandler struct{}

func (h *fakeStateHandler) Help(userID string) string {
	return "help " + userID
}

func (h *fakeStateHandler) NewSession(userID string) string {
	return "new session " + userID
}

func (h *fakeStateHandler) StartMaleRegistration(userID string) string {
	return "start male registration " + userID
}

func (h *fakeStateHandler) EndMaleRegistration(userID string) string {
	return "end male registration " + userID
}

func (h *fakeStateHandler) StartFemaleRegistration(userID string) string {
	return "start female registration " + userID
}

func (h *fakeStateHandler) EndFemaleRegistration(userID string) string {
	return "end female registration " + userID
}

func (h *fakeStateHandler) AddTeamMember(userID string, teamID string) string {
	return "add team member " + userID + " " + teamID
}
