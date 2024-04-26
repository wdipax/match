package state_test

import (
	"testing"

	"github.com/wdipax/match/state"
)

func TestState(t *testing.T) {
	t.Parallel()

	s := state.New()

	var e any

	s.Process(e)
}
