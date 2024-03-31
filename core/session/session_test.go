package session_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/core/player"
	"github.com/wdipax/match/core/session"
	"github.com/wdipax/match/core/team"
)

func TestSession(t *testing.T) {
	t.Parallel()

	t.Run("it allows to add a new team", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New(
			"m", // name
		)

		err := s.AddTeam(tm)

		assert.NoError(t, err)
	})

	t.Run("it does not allow to add the same team twice", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New(
			"m", // name
		)

		err := s.AddTeam(tm)

		require.NoError(t, err)

		err = s.AddTeam(tm)

		assert.Error(t, err, "already existed team was created again")
	})

	t.Run("it computes matching players", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm1 := team.New(
			"m", // name
		)

		require.NoError(t, s.AddTeam(tm1))

		p1 := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // ID
			s,
		)

		require.NoError(t, tm1.AddPlayer(p1))

		tm2 := team.New(
			"f", // name
		)

		require.NoError(t, s.AddTeam(tm2))

		p2 := player.New(
			"@janeroe", // account
			"Jane",     // name
			10,         // ID
			s,
		)

		require.NoError(t, tm2.AddPlayer(p2))

		require.NoError(t, p1.Choose(p2.ID))

		require.NoError(t, p2.Choose(p1.ID))

		assert.ElementsMatch(t, []uint8{p2.ID}, s.PlayerMatches(p1.ID), "wrong match for player1")
		assert.ElementsMatch(t, []uint8{p1.ID}, s.PlayerMatches(p2.ID), "wrong match for player2")
	})
}
