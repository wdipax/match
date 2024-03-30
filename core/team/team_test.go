package team_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/core/team"
	"github.com/wdipax/match/core/user"
)

func TestTeam(t *testing.T) {
	t.Parallel()

	t.Run("it registers a new user", func(t *testing.T) {
		t.Parallel()

		s := team.New("m")

		u := user.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
		)

		err := s.AddUser(u)

		assert.NoError(t, err)
	})

	t.Run("it does not registers already existed user", func(t *testing.T) {
		t.Parallel()

		s := team.New("m")

		u := user.New(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
		)

		err := s.AddUser(u)

		require.NoError(t, err)

		err = s.AddUser(u)

		assert.Error(t, err, "already existed user was registered again")
	})
}
