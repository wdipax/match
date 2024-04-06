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

		var u fakeUpdate

		s.Process(u)
	})
}

type fakeUpdate struct{}

func (fakeUpdate) FromAdmin() bool {
	return false
}

func (fakeUpdate) SendMessage(string) {
	return
}
