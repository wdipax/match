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

	t.Run("it registers a new user", func(t *testing.T) {
		t.Parallel()

		tm := team.New("m")

		s := session.New()

		u := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		err := tm.AddUser(u)

		assert.NoError(t, err)
	})

	t.Run("it does not registers already existed user", func(t *testing.T) {
		t.Parallel()

		tm := team.New("m")

		s := session.New()

		u := player.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
			s,
		)

		err := tm.AddUser(u)

		require.NoError(t, err)

		err = tm.AddUser(u)

		assert.Error(t, err, "already existed user was registered again")
	})
}
