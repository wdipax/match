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

			require.Len(t, res, 1)

			assert.Equal(t, "admin", res[0].UserID)
			assert.Equal(t, "voting has started", res[0].MSG)
		})

		t.Run("all users receive polls", func(t *testing.T) {
			t.Parallel()

			ss, c := stateSettings()

			st := state.New(ss)

			hs := helperSettings{
				state: st,
				core:  c,
			}

			startSession(t, hs)

			hs.teamType = male

			maleTeamID := startTeamRegistration(t, hs)

			joinTeam(t, hs, "male", maleTeamID)

			endTeamRegistration(t, hs)

			hs.teamType = female

			femaleTeamID := startTeamRegistration(t, hs)

			joinTeam(t, hs, "female", femaleTeamID)

			endTeamRegistration(t, hs)

			res := st.StartVoting("admin")

			maleRe := getUserResponse(t, res, "male")

			assert.Equal(t, "please vote for female_team_id", maleRe.MSG)

			femaleRe := getUserResponse(t, res, "female")

			assert.Equal(t, "please vote for male_team_id", femaleRe.MSG)
		})

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

			re := getUserResponse(t, res, "user")

			assert.Equal(t, "vote received", re.MSG)
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

		t.Run("session ends", func(t *testing.T) {
			t.Parallel()

			t.Run("when all users have voted", func(t *testing.T) {
				t.Parallel()

				ss, c := stateSettings()

				st := state.New(ss)

				hs := helperSettings{
					state: st,
					core:  c,
				}

				startSession(t, hs)

				hs.teamType = male

				maleTeamID := startTeamRegistration(t, hs)

				joinTeam(t, hs, "male", maleTeamID)

				endTeamRegistration(t, hs)

				hs.teamType = female

				femaleTeamID := startTeamRegistration(t, hs)

				joinTeam(t, hs, "female", femaleTeamID)

				endTeamRegistration(t, hs)

				st.StartVoting("admin")

				vote(t, hs, "male")

				res := vote(t, hs, "female")

				re := getUserResponse(t, res, "admin")

				assert.Equal(t, "session ended", re.MSG)
			})

			t.Run("when admin ends the session", func(t *testing.T) {
				t.Parallel()

				ss, _ := stateSettings()

				st := state.New(ss)

				res := st.EndSession("admin")

				require.Len(t, res, 1)

				assert.Equal(t, "admin", res[0].UserID)
				assert.Equal(t, "session ended", res[0].MSG)
			})
		})

		t.Run("user can not end the session", func(t *testing.T) {
			t.Parallel()

			ss, _ := stateSettings()

			st := state.New(ss)

			assert.Empty(t, st.EndSession("user"))
		})

		t.Run("users receive their matches", func(t *testing.T) {
			t.Parallel()

			ss, c := stateSettings()

			st := state.New(ss)

			hs := helperSettings{
				state: st,
				core:  c,
			}

			startSession(t, hs)

			hs.teamType = male

			registerUsersToTeam(t, hs, "male1", "male2")

			hs.teamType = female

			registerUsersToTeam(t, hs, "female1", "female2")

			st.StartVoting("admin")

			vote(t, hs, "male1")
			vote(t, hs, "female1")
			vote(t, hs, "male2")
			res := vote(t, hs, "female2")

			getMatches := responseMSGFilter("here is your matches")

			getUserResponse(t, res, "female2", getMatches)
			getUserResponse(t, res, "male1", getMatches)
			getUserResponse(t, res, "male2", getMatches)
			getUserResponse(t, res, "female1", getMatches)
		})

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

func joinTeam(tb testing.TB, settings helperSettings, userID string, teamID string) {
	tb.Helper()

	res := settings.state.Input(userID, teamID)

	require.Len(tb, res, 1)
	require.Equal(tb, userID, res[0].UserID)
	require.Contains(tb, res[0].MSG, "joined team")
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

func registerUsersToTeam(tb testing.TB, settings helperSettings, userIDs ...string) {
	tb.Helper()

	teamID := startTeamRegistration(tb, settings)

	for _, id := range userIDs {
		joinTeam(tb, settings, id, teamID)
	}

	endTeamRegistration(tb, settings)
}

func vote(tb testing.TB, settings helperSettings, userID string) []*state.Response {
	tb.Helper()

	res := settings.state.Input(userID, "vote")

	re := getUserResponse(tb, res, userID)

	require.Equal(tb, "vote received", re.MSG)

	return res
}

func responseMSGFilter(msg string) func(*state.Response) bool {
	return func(r *state.Response) bool {
		return r.MSG == msg
	}
}

func getUserResponse(tb testing.TB, res []*state.Response, userID string, filters ...func(*state.Response) bool) *state.Response {
	tb.Helper()

	var skip bool

	for _, v := range res {
		skip = false

		for _, f := range filters {
			if !f(v) {
				skip = true

				break
			}
		}

		if skip {
			continue
		}

		if v.UserID == userID {
			return v
		}
	}

	tb.Fatalf("no response for %s", userID)

	return nil
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

func (c *fakeCore) Poll(teamID string) string {
	return "please vote for " + teamID
}

func (c *fakeCore) Vote(userID string, voice string) {

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
		StartVotingMSG:         func(optional string) string { return "voting has started" },
		VoteReceivedMSG:        func(optional string) string { return "vote received" },
		EndSessionMSG:          func(optional string) string { return "session ended" },
		MaleTeamName:           "male team",
		FemaleTeamName:         "female team",
	}

	return ss, &core
}
