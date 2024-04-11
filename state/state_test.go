package state_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/state"
)

func TestState(t *testing.T) {
	t.Parallel()

	t.Run("admin can start a new session", func(t *testing.T) {
		t.Parallel()

		var (
			c fakeCore
		)

		a := fakeIsAdmin{
			adminID: "admin",
		}

		st := state.New(a.IsAdmin, &c)

		st.NewSession("admin")

		assert.Equal(t, 1, c.sessions)
	})

	t.Run("user can not start a new session", func(t *testing.T) {
		t.Parallel()

		var (
			c fakeCore
		)

		a := fakeIsAdmin{
			adminID: "admin",
		}

		st := state.New(a.IsAdmin, &c)

		st.NewSession("user")

		assert.Empty(t, c.sessions)
	})
}

type fakeIsAdmin struct {
	adminID string
}

func (a fakeIsAdmin) IsAdmin(userID string) bool {
	return a.adminID == userID
}

type fakeCore struct {
	sessions int
}

func (c *fakeCore) NewSession() string {
	c.sessions++

	return "session_id"
}
