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

		assert.Equal(t, "user", tg.userID)
		assert.Equal(t, "helped", tg.sentMsg)
	})

	t.Run("it starts a new session", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.NewSession,
			userID: "admin",
		}

		e.Process(&evt)

		assert.Equal(t, "admin", tg.userID)
		assert.Equal(t, "started a new session", tg.sentMsg)
	})

	t.Run("it starts male registration", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.StartMaleRegistration,
			userID: "admin",
		}

		e.Process(&evt)

		assert.Equal(t, "admin", tg.userID)
		assert.Equal(t, "started male registration", tg.sentMsg)
	})

	t.Run("it ends male registration", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.EndMaleRegistration,
			userID: "admin",
		}

		e.Process(&evt)

		assert.Equal(t, "admin", tg.userID)
		assert.Equal(t, "ended male registration", tg.sentMsg)
	})

	t.Run("it starts female registration", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.StartFemaleRegistration,
			userID: "admin",
		}

		e.Process(&evt)

		assert.Equal(t, "admin", tg.userID)
		assert.Equal(t, "started female registration", tg.sentMsg)
	})

	t.Run("it ends female registration", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.EndFemaleRegistration,
			userID: "admin",
		}

		e.Process(&evt)

		assert.Equal(t, "admin", tg.userID)
		assert.Equal(t, "ended female registration", tg.sentMsg)
	})

	t.Run("it adds a team member", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action:  engine.AddTeamMember,
			userID:  "user",
			payload: "team",
		}

		e.Process(&evt)

		assert.Equal(t, "user", tg.userID)
		assert.Equal(t, "added a team member for the team", tg.sentMsg)
	})

	t.Run("it sets a team member name", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action:  engine.TeamMemberName,
			userID:  "user",
			payload: "John",
		}

		e.Process(&evt)

		assert.Equal(t, "user", tg.userID)
		assert.Equal(t, "set a team member name to John", tg.sentMsg)
	})

	t.Run("it sets a team member number", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action:  engine.TeamMemberNumber,
			userID:  "user",
			payload: "5",
		}

		e.Process(&evt)

		assert.Equal(t, "user", tg.userID)
		assert.Equal(t, "set a team member number to 5", tg.sentMsg)
	})

	t.Run("it starts voting", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.StartVoting,
			userID: "admin",
		}

		e.Process(&evt)

		assert.Equal(t, "admin", tg.userID)
		assert.Equal(t, "started voting", tg.sentMsg)
	})

	t.Run("it performs voting", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action:  engine.Vote,
			userID:  "user",
			payload: "1,2,3",
		}

		e.Process(&evt)

		assert.Equal(t, "user", tg.userID)
		assert.Equal(t, "voted for 1,2,3", tg.sentMsg)
	})

	t.Run("it ends the session", func(t *testing.T) {
		t.Parallel()

		var (
			tg fakeTelegramHandler
			st fakeStateHandler
		)

		e := engine.New(&tg, &st)

		evt := fakeEvent{
			action: engine.EndSession,
			userID: "admin",
		}

		e.Process(&evt)

		assert.Equal(t, "admin", tg.userID)
		assert.Equal(t, "ended the session", tg.sentMsg)
	})
}

type fakeTelegramHandler struct {
	userID  string
	sentMsg string
}

func (tg *fakeTelegramHandler) Send(userID string, msg string) {
	tg.userID = userID
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
	return "helped"
}

func (h *fakeStateHandler) NewSession(userID string) string {
	return "started a new session"
}

func (h *fakeStateHandler) StartMaleRegistration(userID string) string {
	return "started male registration"
}

func (h *fakeStateHandler) EndMaleRegistration(userID string) string {
	return "ended male registration"
}

func (h *fakeStateHandler) StartFemaleRegistration(userID string) string {
	return "started female registration"
}

func (h *fakeStateHandler) EndFemaleRegistration(userID string) string {
	return "ended female registration"
}

func (h *fakeStateHandler) AddTeamMember(userID string, teamID string) string {
	return "added a team member for the " + teamID
}

func (h *fakeStateHandler) TeamMemberName(userID string, name string) string {
	return "set a team member name to " + name
}

func (h *fakeStateHandler) TeamMemberNumber(userID string, number string) string {
	return "set a team member number to " + number
}

func (h *fakeStateHandler) StartVoting(userID string) string {
	return "started voting"
}

func (h *fakeStateHandler) Vote(userID string, poll string) string {
	return "voted for " + poll
}

func (h *fakeStateHandler) EndSession(userID string) string {
	return "ended the session"
}
