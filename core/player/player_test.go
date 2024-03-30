package player_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/core/player"
)

func TestUser(t *testing.T) {
	t.Parallel()

	t.Run("they can chose another player", func(t *testing.T) {
		t.Parallel()

		p := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
		)

		err := p.Choose(10)

		assert.NoError(t, err)
	})

	t.Run("they can not choose a player from their team", func(t *testing.T) {
		t.Parallel()

		p := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
		)

		err := p.Choose(5)

		assert.Error(t, err, "choosen a player from the same team")
	})
}
