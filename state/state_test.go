package state_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/state"
)

func TestState(t *testing.T) {
	t.Parallel()

	t.Run("it can start a new session", func(t *testing.T) {
		t.Parallel()

		var sh sessionHandler

		st := state.New(&sh)

		st.NewSession()

		assert.Len(t, sh.sessions, 1)
	})
}

type sessionHandler struct {
	sessions []int
}

func (h *sessionHandler) New() {
	h.sessions = append(h.sessions, 1)
}
