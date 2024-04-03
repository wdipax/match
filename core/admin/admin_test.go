package admin_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wdipax/match/core/admin"
)

func TestAdmin(t *testing.T) {
	t.Parallel()

	t.Run("they can start a new session", func(t *testing.T) {
		t.Parallel()

		a := admin.New()

		err := a.NewSession()

		assert.NoError(t, err)
	})

	t.Run("they can end the session", func(t *testing.T) {
		t.Parallel()

		a := admin.New()

		require.NoError(t, a.NewSession())

		assert.NoError(t, a.EndSession())
	})
}
