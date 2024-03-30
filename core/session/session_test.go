package session_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/core/session"
)

func TestSession(t *testing.T) {
	t.Parallel()

	t.Run("it allows to create a new team", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		err := s.AddTeam(
			"m", // name
		)

		assert.NoError(t, err)
	})

	t.Run("it does not allow to create the same team twice", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		err := s.AddTeam(
			"m", // name
		)

		require.NoError(t, err)

		err = s.AddTeam(
			"m", // name
		)

		assert.Error(t, err, "already existed team was created again")
	})
}
