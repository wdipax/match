package state_test

import (
	"testing"

	"github.com/wdipax/match/state"
)

func TestState(t *testing.T) {
	t.Parallel()

	t.Run("it processes updates", func(t *testing.T) {
		t.Parallel()

		s := state.New()

		s.Process(nil)
	})
}
