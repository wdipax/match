package state_test

import (
	"slices"
	"testing"

	"github.com/google/uuid"
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

	t.Run("it can end a session", func(t *testing.T) {
		t.Parallel()

		var sh sessionHandler

		st := state.New(&sh)

		id := st.NewSession()

		st.EndSession(id)

		assert.Empty(t, sh.sessions)
	})
}

type sessionHandler struct {
	sessions []string
}

func (h *sessionHandler) New() string {
	id := uuid.NewString()

	h.sessions = append(h.sessions, id)

	return id
}

func (h *sessionHandler) Delete(id string) {
	for i := range h.sessions {
		if h.sessions[i] == id {
			h.sessions = slices.Delete(h.sessions, i, i+1)

			return
		}
	}
}
