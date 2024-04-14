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

			st := state.New(state.StateSettings{
				IsAdmin: a.IsAdmin,
				Core:    &c,
			})

			st.NewSession("admin")

			assert.Equal(t, 1, c.sessions)
		})

		t.Run("user can not start a new session", func(t *testing.T) {
			t.Parallel()

			var c fakeCore

			a := fakeIsAdmin{
				adminID: "admin",
			}

			st := state.New(state.StateSettings{
				IsAdmin: a.IsAdmin,
				Core:    &c,
			})

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

				st := state.New(state.StateSettings{
					IsAdmin: a.IsAdmin,
					Core:    &c,
				})

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

				st := state.New(state.StateSettings{
					IsAdmin: a.IsAdmin,
					Core:    &c,
				})

				res := st.StartFemaleRegistration("admin")

				require.Len(t, res, 1)

				assert.Equal(t, "admin", res[0].UserID)
				assert.Equal(t, "female_team_id", res[0].MSG)
			})
		})

		t.Run("user can join a team", func(t *testing.T) {
			t.Parallel()

			t.Run("male team", func(t *testing.T) {
				t.Parallel()

				var c fakeCore

				a := fakeIsAdmin{
					adminID: "admin",
				}

				st := state.New(state.StateSettings{
					IsAdmin:      a.IsAdmin,
					JoinTeamMSG:  joinTeamMSG,
					Core:         &c,
					MaleTeamName: "male team",
				})

				teamID := startTeamRegistration(t, helperSettings{
					state:    st,
					teamType: male,
					core:     &c,
					adminID:  a.adminID,
				})

				res := st.Input("user", teamID)

				require.Len(t, res, 1)

				assert.Equal(t, "user", res[0].UserID)
				assert.Contains(t, res[0].MSG, "male team")
			})

			t.Run("female team", func(t *testing.T) {
				t.Parallel()

				var c fakeCore

				a := fakeIsAdmin{
					adminID: "admin",
				}

				st := state.New(state.StateSettings{
					IsAdmin:        a.IsAdmin,
					JoinTeamMSG:    joinTeamMSG,
					Core:           &c,
					FemaleTeamName: "female team",
				})

				teamID := startTeamRegistration(t, helperSettings{
					state:    st,
					teamType: female,
					core:     &c,
					adminID:  a.adminID,
				})

				res := st.Input("user", teamID)

				require.Len(t, res, 1)

				assert.Equal(t, "user", res[0].UserID)
				assert.Contains(t, res[0].MSG, "female team")
			})
		})

		t.Run("admin can end registration to a team", func(t *testing.T) {
			t.Parallel()

			t.Run("male team", func(t *testing.T) {
				t.Parallel()

				var c fakeCore

				a := fakeIsAdmin{
					adminID: "admin",
				}

				st := state.New(state.StateSettings{
					IsAdmin:      a.IsAdmin,
					JoinTeamMSG:  joinTeamMSG,
					Core:         &c,
					MaleTeamName: "male team",
				})

				startTeamRegistration(t, helperSettings{
					state:    st,
					teamType: male,
					core:     &c,
					adminID:  a.adminID,
				})

				res := st.EndMaleRegistration("admin")

				require.Len(t, res, 1)

				assert.Equal(t, "admin", res[0].UserID)
				assert.Contains(t, res[0].MSG, "male team")
			})

			t.Run("female team", func(t *testing.T) {
				t.Parallel()

				
			})
		})
	})
}

type teamType int

const (
	male teamType = iota
	female
)

type helperSettings struct {
	state    *state.State
	teamType teamType
	core     *fakeCore
	adminID  string
}

func startTeamRegistration(tb testing.TB, settings helperSettings) string {
	tb.Helper()

	defer func(original string) {
		settings.core.newTeamID = original
	}(settings.core.newTeamID)

	var startTeamRegistration func(userID string) []*state.Response

	switch settings.teamType {
	case male:
		settings.core.newTeamID = "male_team_id"

		startTeamRegistration = settings.state.StartMaleRegistration
	case female:
		settings.core.newTeamID = "female_team_id"

		startTeamRegistration = settings.state.StartFemaleRegistration
	}

	res := startTeamRegistration(settings.adminID)

	require.Len(tb, res, 1)

	return res[0].MSG
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

func (c *fakeCore) NewTeam(name string) string {
	return c.newTeamID
}

func joinTeamMSG(teamName string) string {
	return "you joined the team " + teamName
}
