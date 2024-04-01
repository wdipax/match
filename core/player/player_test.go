package player_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

		s.AddTeam(tm1)

		p1 := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		tm1.AddPlayer(p1)

		tm2 := team.New("f", s)

		s.AddTeam(tm2)

		p2 := player.New(
			"@pineapple", // account
			"Alice",      // name
			10,           // id
			s,
		)

		tm2.AddPlayer(p2)

		err := p1.Choose(10)

		assert.NoError(t, err)
	})

	t.Run("they can not choose a player from their team", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New("m", s)

		s.AddTeam(tm)

		p := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		tm.AddPlayer(p)

		err := p.Choose(5)

		assert.Error(t, err, "choosen a player from the same team")
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
