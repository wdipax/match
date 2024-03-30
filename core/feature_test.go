package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/core"
)

func TestRegistration(t *testing.T) {
	t.Parallel()

	t.Run("it registers a new user", func(t *testing.T) {
		t.Parallel()

		err := core.Register(
			"@raspberry", // account
			"Dima",       // name
			5,            // id
		)

		assert.NoError(t, err)
	})
}
