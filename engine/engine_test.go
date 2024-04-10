package engine_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/engine"
	"github.com/wdipax/match/state"
)

func TestEngine(t *testing.T) {
	t.Parallel()

	t.Run("it responds with help message", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.Help,
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

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.NewSession,
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

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.StartMaleRegistration,
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

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.EndMaleRegistration,
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

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.StartFemaleRegistration,
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

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.EndFemaleRegistration,
			userID: "admin",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "admin",
			MSG:    "ended female registration",
		}}, tg.sent)
	})

	t.Run("it adds a team member", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action:  engine.AddTeamMember,
			userID:  "user",
			payload: "team",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "user",
			MSG:    "added as a team member for the team",
		}}, tg.sent)
	})

	t.Run("it sets a team member name", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action:  engine.TeamMemberName,
			userID:  "user",
			payload: "John",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "user",
			MSG:    "set a team member name to John",
		}}, tg.sent)
	})

	t.Run("it sets a team member number", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action:  engine.TeamMemberNumber,
			userID:  "user",
			payload: "5",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "user",
			MSG:    "set a team member number to 5",
		}}, tg.sent)
	})

	t.Run("it starts voting", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.StartVoting,
			userID: "admin",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "admin",
			MSG:    "started voting",
		}}, tg.sent)
	})

	t.Run("it performs voting", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action:  engine.Vote,
			userID:  "user",
			payload: "1,2,3",
		}

		e.Process(&evt)

		assert.ElementsMatch(t, []*state.Response{{
			UserID: "user",
			MSG:    "voted for 1,2,3",
		}}, tg.sent)
	})

	t.Run("it ends the session", func(t *testing.T) {
		t.Parallel()

		var (
			tg spyMessengerHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.EndSession,
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

func (h *fakeStateHandler) TeamMemberName(userID string, name string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "set a team member name to " + name,
	}}
}

func (h *fakeStateHandler) TeamMemberNumber(userID string, number string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "set a team member number to " + number,
	}}
}

func (h *fakeStateHandler) StartVoting(userID string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "started voting",
	}}
}

func (h *fakeStateHandler) Vote(userID string, poll string) []*state.Response {
	return []*state.Response{{
		UserID: userID,
		MSG:    "voted for " + poll,
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
