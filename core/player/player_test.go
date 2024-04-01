package player_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/core/player"
	"github.com/wdipax/match/core/session"
	"github.com/wdipax/match/core/team"
)

func TestUser(t *testing.T) {
	t.Parallel()

	t.Run("they can chose another player", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm1 := team.New("m", s)

		require.NoError(t, s.AddTeam(tm1))

		p1 := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		require.NoError(t, tm1.AddPlayer(p1))

		tm2 := team.New("f", s)

		require.NoError(t, s.AddTeam(tm2))

		p2 := player.New(
			"@pineapple", // account
			"Alice",      // name
			5,           // id
			s,
		)

		require.NoError(t, tm2.AddPlayer(p2))

		assert.NoError(t, p1.Choose(p2.ID))
	})

	t.Run("they can not choose a player from their team", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New("m", s)

		require.NoError(t, s.AddTeam(tm))

		p1 := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		p2 := player.New(
			"@pineapple", // account
			"Alice",      // name
			6,            // id
			s,
		)

		require.NoError(t, tm.AddPlayer(p1))

		require.NoError(t, tm.AddPlayer(p2))

		assert.Error(t, p1.Choose(p2.ID), "choosen a player from the same team")
	})

	t.Run("it compares players by account", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		p1 := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		p2 := player.New(
			"@raspberry", // account
			"John",       // name
			55,           // id
			s,
		)

		assert.True(t, player.TheSame(p1, p2), "should be the same player")
	})
}
