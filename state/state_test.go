package state_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/state"
)

func TestState(t *testing.T) {
	t.Parallel()

	t.Run("session", func(t *testing.T) {
		t.Parallel()

		t.Run("admin can start a new session", func(t *testing.T) {
			t.Parallel()

			var c fakeCore

			a := fakeIsAdmin{
				adminID: "admin",
			}

			st := state.New(a.IsAdmin, &c)

			st.NewSession("admin")

			assert.Equal(t, 1, c.sessions)
		})

		t.Run("user can not start a new session", func(t *testing.T) {
			t.Parallel()

			var c fakeCore

			a := fakeIsAdmin{
				adminID: "admin",
			}

			st := state.New(a.IsAdmin, &c)

			st.NewSession("user")

			assert.Empty(t, c.sessions)
		})
	})

	t.Run("team", func(t *testing.T) {
		t.Run("admin can start registration to a team", func(t *testing.T) {
			t.Parallel()

			t.Run("male team", func(t *testing.T) {
				t.Parallel()

				c := fakeCore{
					newTeamID: "male_team_id",
				}

				a := fakeIsAdmin{
					adminID: "admin",
				}

				st := state.New(a.IsAdmin, &c)

				res := st.StartMaleRegistration("admin")

				require.Len(t, res, 1)

				assert.Equal(t, "admin", res[0].UserID)
				assert.Equal(t, "male_team_id", res[0].MSG)
			})

			t.Run("female team", func(t *testing.T) {
				t.Parallel()

				c := fakeCore{
					newTeamID: "female_team_id",
				}

				a := fakeIsAdmin{
					adminID: "admin",
				}

				st := state.New(a.IsAdmin, &c)

				res := st.StartFemaleRegistration("admin")

				require.Len(t, res, 1)

				assert.Equal(t, "admin", res[0].UserID)
				assert.Equal(t, "female_team_id", res[0].MSG)
			})
		})
	})
}

type fakeIsAdmin struct {
	adminID string
}

func (a fakeIsAdmin) IsAdmin(userID string) bool {
	return a.adminID == userID
}

type fakeCore struct {
	newTeamID string

	sessions int
}

func (c *fakeCore) NewSession() string {
	c.sessions++

	return "session_id"
}

func (c *fakeCore) NewTeam() string {
	return c.newTeamID
}
