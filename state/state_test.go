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

			ss, _ := stateSettings()

			st := state.New(ss)

			res := st.NewSession("admin")

			require.Len(t, res, 1)

			assert.Equal(t, "admin", res[0].UserID)
			assert.Equal(t, "session started", res[0].MSG)
		})

		t.Run("user can not start a new session", func(t *testing.T) {
			t.Parallel()

			ss, _ := stateSettings()

			st := state.New(ss)

			assert.Empty(t, st.NewSession("user"))
		})
	})

	t.Run("team", func(t *testing.T) {
		t.Run("admin can start registration to a team", func(t *testing.T) {
			t.Parallel()

			t.Run("male team", func(t *testing.T) {
				t.Parallel()

				ss, c := stateSettings()

				c.newTeamID = "male_team_id"

				st := state.New(ss)

				startSession(t, helperSettings{
					state: st,
				})

				res := st.StartMaleRegistration("admin")

				require.Len(t, res, 1)

				assert.Equal(t, "admin", res[0].UserID)
				assert.Equal(t, "male_team_id", res[0].MSG)
			})

			t.Run("female team", func(t *testing.T) {
				t.Parallel()

				ss, c := stateSettings()

				c.newTeamID = "female_team_id"

				st := state.New(ss)

				startSession(t, helperSettings{
					state: st,
				})

				res := st.StartFemaleRegistration("admin")

				require.Len(t, res, 1)

				assert.Equal(t, "admin", res[0].UserID)
				assert.Equal(t, "female_team_id", res[0].MSG)
			})
		})

		t.Run("user can not start registration to a team", func(t *testing.T) {
			t.Parallel()

			t.Run("male team", func(t *testing.T) {
				t.Parallel()

				ss, _ := stateSettings()

				st := state.New(ss)

				startSession(t, helperSettings{
					state: st,
				})

				assert.Empty(t, st.StartMaleRegistration("user"))
			})

			t.Run("female team", func(t *testing.T) {
				t.Parallel()

				ss, _ := stateSettings()

				st := state.New(ss)

				startSession(t, helperSettings{
					state: st,
				})

				assert.Empty(t, st.StartFemaleRegistration("user"))
			})
		})

		t.Run("user can join a team", func(t *testing.T) {
			t.Parallel()

			t.Run("male team", func(t *testing.T) {
				t.Parallel()

				ss, c := stateSettings()

				st := state.New(ss)

				hs := helperSettings{
					state:    st,
					teamType: male,
					core:     c,
				}

				startSession(t, hs)

				teamID := startTeamRegistration(t, hs)

				res := st.Input("user", teamID)

				require.Len(t, res, 1)

				assert.Equal(t, "user", res[0].UserID)
				assert.Contains(t, res[0].MSG, "male team")
			})

			t.Run("female team", func(t *testing.T) {
				t.Parallel()

				ss, c := stateSettings()

				st := state.New(ss)

				hs := helperSettings{
					state:    st,
					teamType: female,
					core:     c,
				}

				startSession(t, hs)

				teamID := startTeamRegistration(t, hs)

				res := st.Input("user", teamID)

				require.Len(t, res, 1)

				assert.Equal(t, "user", res[0].UserID)
				assert.Contains(t, res[0].MSG, "female team")
			})
		})

		t.Run("admin can not join a team", func(t *testing.T) {
			t.Parallel()

			ss, c := stateSettings()

			st := state.New(ss)

			hs := helperSettings{
				state:    st,
				teamType: male,
				core:     c,
			}

			startSession(t, hs)

			teamID := startTeamRegistration(t, hs)

			res := st.Input("admin", teamID)

			require.Len(t, res, 1)

			assert.Equal(t, "admin", res[0].UserID)
			assert.Equal(t, "admin can not join a team", res[0].MSG)
		})

		t.Run("admin can end registration to a team", func(t *testing.T) {
			t.Parallel()

			t.Run("male team", func(t *testing.T) {
				t.Parallel()

				ss, c := stateSettings()

				st := state.New(ss)

				hs := helperSettings{
					state:    st,
					teamType: male,
					core:     c,
				}

				startSession(t, hs)

				startTeamRegistration(t, hs)

				res := st.EndMaleRegistration("admin")

				require.Len(t, res, 1)

				assert.Equal(t, "admin", res[0].UserID)
				assert.Contains(t, res[0].MSG, "male team")
			})

			t.Run("female team", func(t *testing.T) {
				t.Parallel()

				ss, c := stateSettings()

				st := state.New(ss)

				hs := helperSettings{
					state:    st,
					teamType: female,
					core:     c,
				}

				startSession(t, hs)

				startTeamRegistration(t, hs)

				res := st.EndFemaleRegistration("admin")

				require.Len(t, res, 1)

				assert.Equal(t, "admin", res[0].UserID)
				assert.Contains(t, res[0].MSG, "female team")
			})
		})

		t.Run("user can not end registration to a team", func(t *testing.T) {
			t.Parallel()

			t.Run("male team", func(t *testing.T) {
				t.Parallel()

				ss, c := stateSettings()

				st := state.New(ss)

				hs := helperSettings{
					state:    st,
					teamType: male,
					core:     c,
				}

				startSession(t, hs)

				startTeamRegistration(t, hs)

				assert.Empty(t, st.EndMaleRegistration("user"))
			})

			t.Run("female team", func(t *testing.T) {
				t.Parallel()

				ss, c := stateSettings()

				st := state.New(ss)

				hs := helperSettings{
					state:    st,
					teamType: female,
					core:     c,
				}

				startSession(t, hs)

				startTeamRegistration(t, hs)

				assert.Empty(t, st.EndFemaleRegistration("user"))
			})
		})

		t.Run("user can not join a team", func(t *testing.T) {
			t.Parallel()

			t.Run("before the registration started", func(t *testing.T) {
				t.Parallel()

				t.Run("session is just created", func(t *testing.T) {
					t.Parallel()

					ss, _ := stateSettings()

					st := state.New(ss)

					startSession(t, helperSettings{
						state: st,
					})

					assert.Empty(t, st.Input("user", "team_id"))
				})

				t.Run("female started, male not", func(t *testing.T) {
					t.Parallel()

					ss, c := stateSettings()

					st := state.New(ss)

					hs := helperSettings{
						state:    st,
						teamType: female,
						core:     c,
					}

					startSession(t, hs)

					startTeamRegistration(t, hs)

					assert.Empty(t, st.Input("user", "male_team_id"))
				})

				t.Run("male started, female not", func(t *testing.T) {
					t.Parallel()

					ss, c := stateSettings()

					st := state.New(ss)

					hs := helperSettings{
						state:    st,
						teamType: male,
						core:     c,
					}

					startSession(t, hs)

					startTeamRegistration(t, hs)

					assert.Empty(t, st.Input("user", "female_team_id"))
				})
			})

			t.Run("after the registration ended", func(t *testing.T) {
				t.Parallel()

				t.Run("male", func(t *testing.T) {
					t.Parallel()

					ss, c := stateSettings()

					st := state.New(ss)

					hs := helperSettings{
						state:    st,
						teamType: male,
						core:     c,
					}

					startSession(t, hs)

					teamID := startTeamRegistration(t, hs)

					endTeamRegistration(t, hs)

					assert.Empty(t, st.Input("user", teamID))
				})

				t.Run("female", func(t *testing.T) {
					t.Parallel()

					ss, c := stateSettings()

					st := state.New(ss)

					hs := helperSettings{
						state:    st,
						teamType: female,
						core:     c,
					}

					startSession(t, hs)

					teamID := startTeamRegistration(t, hs)

					endTeamRegistration(t, hs)

					assert.Empty(t, st.Input("user", teamID))
				})
			})
		})
	})

	t.Run("voting", func(t *testing.T) {
		t.Parallel()

		t.Run("admin can start voting", func(t *testing.T) {
			t.Parallel()

			ss, _ := stateSettings()

			st := state.New(ss)

			startSession(t, helperSettings{
				state: st,
			})

			res := st.StartVoting("admin")

			// TODO: should receive confirmation.
			assert.NotEmpty(t, res)
		})

		// t.Run("all users receive polls", func(t *testing.T) {
		// 	t.Parallel()

		// 	var c fakeCore

		// 	a := fakeIsAdmin{
		// 		adminID: "admin",
		// 	}

		// 	st := state.New(state.StateSettings{
		// 		IsAdmin:     a.IsAdmin,
		// 		JoinTeamMSG: joinTeamMSG,
		// 		Core:        &c,
		// 	})

		// })

		t.Run("user can not start voting", func(t *testing.T) {
			t.Parallel()

			ss, _ := stateSettings()

			st := state.New(ss)

			assert.Empty(t, st.StartVoting("user"))
		})

		t.Run("user can vote", func(t *testing.T) {
			t.Parallel()

			ss, c := stateSettings()

			st := state.New(ss)

			hs := helperSettings{
				state:    st,
				teamType: male,
				core:     c,
			}

			startSession(t, hs)

			teamID := startTeamRegistration(t, hs)

			st.Input("user", teamID)

			res := st.StartVoting("admin")

			require.NotEmpty(t, res)

			res = st.Input("user", "vote")

			require.Len(t, res, 1)

			assert.Equal(t, "user", res[0].UserID)
			assert.Equal(t, "vote received", res[0].MSG)
		})

		t.Run("amdin can not vote", func(t *testing.T) {
			t.Parallel()

			ss, _ := stateSettings()

			st := state.New(ss)

			hs := helperSettings{
				state: st,
			}

			startSession(t, hs)

			st.StartVoting("admin")

			res := st.Input("admin", "vote")

			require.Len(t, res, 1)

			assert.Equal(t, "admin", res[0].UserID)
			assert.Equal(t, "admin can not vote", res[0].MSG)
		})

		// TODO: voting ends
		// TODO: when all user has voted

		t.Run("when admin ends the session", func(t *testing.T) {
			t.Parallel()

			ss, _ := stateSettings()

			st := state.New(ss)

			res := st.EndSession("admin")

			require.Len(t, res, 1)

			assert.Equal(t, "admin", res[0].UserID)
			assert.Equal(t, "session ended", res[0].MSG)
		})

		t.Run("user can not end the session", func(t *testing.T) {
			t.Parallel()

			ss, _ := stateSettings()

			st := state.New(ss)

			assert.Empty(t, st.EndSession("user"))
		})

		// TODO: users receive their matches

		// TODO: admin does not receive matches
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
}

func startSession(tb testing.TB, settings helperSettings) {
	tb.Helper()

	res := settings.state.NewSession("admin")

	require.Len(tb, res, 1)
	require.Equal(tb, "admin", res[0].UserID)
	require.Equal(tb, "session started", res[0].MSG)
}

func startTeamRegistration(tb testing.TB, settings helperSettings) string {
	tb.Helper()

	var startTeamRegistration func(userID string) []*state.Response

	switch settings.teamType {
	case male:
		settings.core.newTeamID = "male_team_id"

		startTeamRegistration = settings.state.StartMaleRegistration
	case female:
		settings.core.newTeamID = "female_team_id"

		startTeamRegistration = settings.state.StartFemaleRegistration
	}

	res := startTeamRegistration("admin")

	require.Len(tb, res, 1)

	return res[0].MSG
}

func endTeamRegistration(tb testing.TB, settings helperSettings) {
	tb.Helper()

	var (
		endTeamRegistration func(userID string) []*state.Response
		msg                 string
	)

	switch settings.teamType {
	case male:
		endTeamRegistration = settings.state.EndMaleRegistration

		msg = "male team registration ended"
	case female:
		endTeamRegistration = settings.state.EndFemaleRegistration

		msg = "female team registration ended"
	}

	res := endTeamRegistration("admin")

	require.Len(tb, res, 1)
	require.Equal(tb, "admin", res[0].UserID)
	require.Equal(tb, msg, res[0].MSG)
}

type fakeCore struct {
	newTeamID string
}

func (c *fakeCore) NewSession() string {
	return "session_id"
}

func (c *fakeCore) NewTeam(name string) string {
	return c.newTeamID
}

func stateSettings() (state.StateSettings, *fakeCore) {
	var core fakeCore

	ss := state.StateSettings{
		IsAdmin:                func(userID string) bool { return userID == "admin" },
		Core:                   &core,
		StartSessionMSG:        func(optional string) string { return "session started" },
		EndTeamMSG:             func(optional string) string { return optional + " registration ended" },
		AdminCanNotJoinTeamMSG: func(optional string) string { return "admin can not join a team" },
		JoinTeamMSG:            func(optional string) string { return "you joined team: " + optional },
		AdminCanNotVoteMSG:     func(optional string) string { return "admin can not vote" },
		VoteReceivedMSG:        func(optional string) string { return "vote received" },
		EndSessionMSG:          func(optional string) string { return "session ended" },
		MaleTeamName:           "male team",
		FemaleTeamName:         "female team",
	}

	return ss, &core
}
