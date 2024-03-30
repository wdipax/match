package player_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/core/player"
	"github.com/wdipax/match/core/session"
)

func TestUser(t *testing.T) {
	t.Parallel()

	t.Run("they can chose another player", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		p := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		err := p.Choose(10)

		assert.NoError(t, err)
	})

	t.Run("they can not choose a player from their team", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		p := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		err := p.Choose(5)

		assert.Error(t, err, "choosen a player from the same team")
	})
}
