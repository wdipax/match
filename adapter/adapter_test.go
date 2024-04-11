package adapter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/adapter"
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/state"
)

func TestEngine(t *testing.T) {
	t.Parallel()

	t.Run("it accepts input", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action:  event.Input,
			userID:  "user",
			payload: "some text",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "user",
			MSG:    "some text",
		}}, tg.sent)
	})

	t.Run("it responds with help message", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action: event.Help,
			userID: "user",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "user",
			MSG:    "helped",
		}}, tg.sent)
	})

	t.Run("it starts a new session", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action: event.NewSession,
			userID: "admin",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "admin",
			MSG:    "started a new session",
		}}, tg.sent)
	})

	t.Run("it starts male registration", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action: event.StartMaleRegistration,
			userID: "admin",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "admin",
			MSG:    "started male registration",
		}}, tg.sent)
	})

	t.Run("it ends male registration", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action: event.EndMaleRegistration,
			userID: "admin",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "admin",
			MSG:    "ended male registration",
		}}, tg.sent)
	})

	t.Run("it starts female registration", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action: event.StartFemaleRegistration,
			userID: "admin",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "admin",
			MSG:    "started female registration",
		}}, tg.sent)
	})

	t.Run("it ends female registration", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action: event.EndFemaleRegistration,
			userID: "admin",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "admin",
			MSG:    "ended female registration",
		}}, tg.sent)
	})

	t.Run("it changes user name", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action: event.ChangeUserName,
			userID: "user",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "user",
			MSG:    "enter your name",
		}}, tg.sent)
	})

	t.Run("it changes user number", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action:  event.ChangeUserNumber,
			userID:  "user",
			payload: "5",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "user",
			MSG:    "enter your number",
		}}, tg.sent)
	})

	t.Run("it starts voting", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action: event.StartVoting,
			userID: "admin",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "admin",
			MSG:    "started voting",
		}}, tg.sent)
	})

	t.Run("it ends the session", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := adapter.New(&tg, &st)

		evt := fakeEvent{
			action: event.EndSession,
			userID: "admin",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "admin",
			MSG:    "ended the session",
		}, {
			UserID: "John",
			MSG:    "received his match Jane",
		}, {
			UserID: "Jane",
			MSG:    "received his match John",
		}}, tg.sent)
	})
}

type spyMessengerHandler struct {
	sent []*state.Response
}

func (tg *spyMessengerHandler) Send(userID string, msg string) {
	tg.sent = append(tg.sent, &state.Response{
		UserID: userID,
		MSG:    msg,
	})
}

type fakeEvent struct {
	action  event.Type
	userID  string
	payload string
}

func (e *fakeEvent) Command() event.Type {
	return e.action
}

func (e *fakeEvent) UserID() string {
	return e.userID
}

func (e *fakeEvent) Payload() string {
	return e.payload
}

type fakeStateHandler struct{}

func (h *fakeStateHandler) Input(userID string, payload string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    payload,
	}}
}

func (h *fakeStateHandler) Help(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "helped",
	}}
}

func (h *fakeStateHandler) NewSession(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "started a new session",
	}}
}

func (h *fakeStateHandler) StartMaleRegistration(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "started male registration",
	}}
}

func (h *fakeStateHandler) EndMaleRegistration(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "ended male registration",
	}}
}

func (h *fakeStateHandler) StartFemaleRegistration(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "started female registration",
	}}
}

func (h *fakeStateHandler) EndFemaleRegistration(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "ended female registration",
	}}
}

func (h *fakeStateHandler) AddTeamMember(userID string, teamID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "added as a team member for the " + teamID,
	}}
}

func (h *fakeStateHandler) ChangeUserName(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "enter your name",
	}}
}

func (h *fakeStateHandler) ChangeUserNumber(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "enter your number",
	}}
}

func (h *fakeStateHandler) StartVoting(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "started voting",
	}}
}

func (h *fakeStateHandler) EndSession(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "ended the session",
	}, {
		UserID: "John",
		MSG:    "received his match Jane",
	}, {
		UserID: "Jane",
		MSG:    "received his match John",
	}}
}
