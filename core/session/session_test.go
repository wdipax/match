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
			s,
		)

		err := s.AddTeam(tm)

		assert.NoError(t, err)
	})

	t.Run("it does not allow to add the same team twice", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New(
			"m", // name
			s,
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
			s,
		)

		require.NoError(t, s.AddTeam(tm1))

		m1 := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // ID
			s,
		)

		m2 := player.New(
			"@johndoe", // account
			"John",     // name
			1,          // ID
			s,
		)

		require.NoError(t, tm1.AddPlayer(m1))
		require.NoError(t, tm1.AddPlayer(m2))

		tm2 := team.New(
			"f", // name
			s,
		)

		require.NoError(t, s.AddTeam(tm2))

		f1 := player.New(
			"@janeroe", // account
			"Jane",     // name
			10,         // ID
			s,
		)

		f2 := player.New(
			"@pineapple", // account
			"Alice",      // name
			2,            // ID
			s,
		)

		require.NoError(t, tm2.AddPlayer(f1))
		require.NoError(t, tm2.AddPlayer(f2))

		require.NoError(t, m1.Choose(f1.ID))
		require.NoError(t, m1.Choose(f2.ID))

		require.NoError(t, m2.Choose(f1.ID))

		require.NoError(t, f1.Choose(m1.ID))
		require.NoError(t, f1.Choose(m2.ID))

		assert.ElementsMatch(t, []uint8{f1.ID}, s.PlayerMatches(m1.ID), "wrong match for m1")

		assert.ElementsMatch(t, []uint8{f1.ID}, s.PlayerMatches(m2.ID), "wrong match for m2")

		assert.ElementsMatch(t, []uint8{m1.ID, m2.ID}, s.PlayerMatches(f1.ID), "wrong match for f1")

		assert.Empty(t, s.PlayerMatches(f2.ID), "wrong match for f2")
	})
}
