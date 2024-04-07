package engine_test

import (
	"testing"

	"github.com/wdipax/match/engine"
)

func TestEngine(t *testing.T) {
	t.Parallel()

	t.Run("it starts a new session", func(t *testing.T) {
		t.Parallel()

		e := engine.New()

		var u fakeEvent

		e.Process(u)
	})
}

type fakeEvent struct{}

func (fakeEvent) FromAdmin() bool {
	return false
}

func (fakeEvent) SendMessage(string) {}
