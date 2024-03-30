package session_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/session"
)

func TestTeam(t *testing.T) {
	t.Parallel()

	t.Run("it allows to create a new team", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		err := s.NewTeam(
			"m", // name
		)

		assert.NoError(t, err)
	})

	t.Run("it does not allow to create the same team twice", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		err := s.NewTeam(
			"m", // name
		)

		require.NoError(t, err)

		err = s.NewTeam(
			"m", // name
		)

		assert.Error(t, err, "already existed team was created again")
	})
}

func TestUserRegistration(t *testing.T) {
	t.Parallel()

	t.Run("it registers a new user", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		err := s.Register(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
		)

		assert.NoError(t, err)
	})

	t.Run("it does not registers already existed user", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		err := s.Register(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
		)

		require.NoError(t, err)

		err = s.Register(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
		)

		assert.Error(t, err, "already existed user was registered again")
	})
}

func TestChoosing(t *testing.T) {
	t.Parallel()

	t.Run("it allows to chose a player from another team", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		err := s.Choose(5)

		assert.NoError(t, err)
	})
}
