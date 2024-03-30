package state_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/state"
)

func TestRegistration(t *testing.T) {
	t.Parallel()

	t.Run("it registers a new user", func(t *testing.T) {
		t.Parallel()

		s := state.New()

		err := s.Register(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
		)

		assert.NoError(t, err)
	})

	t.Run("it does not registers already existed user", func(t *testing.T) {
		t.Parallel()

		s := state.New()

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
