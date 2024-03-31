package session_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/core/session"
	"github.com/wdipax/match/core/team"
)

func TestSession(t *testing.T) {
	t.Parallel()

	t.Run("it allows to add a new team", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New(
			"m", // name
		)

		err := s.AddTeam(tm)

		assert.NoError(t, err)
	})

	t.Run("it does not allow to add the same team twice", func(t *testing.T) {
		t.Parallel()

		s := session.New()

		tm := team.New(
			"m", // name
		)

		err := s.AddTeam(tm)

		require.NoError(t, err)

		err = s.AddTeam(tm)

		assert.Error(t, err, "already existed team was created again")
	})
}
