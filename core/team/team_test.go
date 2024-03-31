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

		tm := team.New("m")

		s := session.New()

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

		tm := team.New("m")

		s := session.New()

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

		tm := team.New("m")

		s := session.New()

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
}
