package team_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/core/player"
	"github.com/wdipax/match/core/session"
	"github.com/wdipax/match/core/team"
)

func TestTeam(t *testing.T) {
	t.Parallel()

	t.Run("it registers a new player", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New("m", s)

		p := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		err := tm.AddPlayer(p)

		assert.NoError(t, err)
	})

	t.Run("it does not registers already existed player", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New("m", s)

		p := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		err := tm.AddPlayer(p)

		require.NoError(t, err)

		err = tm.AddPlayer(p)

		assert.Error(t, err, "already existed player was registered again")
	})

	t.Run("it does not registers players with the same id", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New("m", s)

		p1 := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		err := tm.AddPlayer(p1)

		require.NoError(t, err)

		p2 := player.New(
			"@johndoe", // account
			"John",     // name
			5,          // id
			s,
		)

		err = tm.AddPlayer(p2)

		assert.Error(t, err, "a player with the same id was registered")
	})

	t.Run("it does not registers a player that already registered in another team", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm1 := team.New("m", s)

		s.AddTeam(tm1)

		p1 := player.New(
			"@sneaky", // account
			"Tricky",  // name
			42,        // id
			s,
		)

		err := tm1.AddPlayer(p1)

		require.NoError(t, err)

		tm2 := team.New("f", s)

		s.AddTeam(tm2)

		p2 := player.New(
			"@sneaky", // account
			"Susan",   // name
			24,        // id
			s,
		)

		err = tm2.AddPlayer(p2)

		assert.Error(t, err, "the spy was regiseterd in both teams!")
	})
}
