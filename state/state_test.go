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

	// The game.
	// Add the Male team.
	// Open registration to the Male team.
	// Add members to the Male team.
	// Close registration to the Male team.
	// Add Female team.
	// Open registration to the Female team.
	// Add members to the Female team.
	// Close registration to the Female team.
	// Know each other.
	// Vote for each other.
	// Receive matches.
	// Delete data.

	// Actions.
	// Start/End a session. [admin]
	// CRUD Team. [admin]
	// Open/close registration to the team. [admin]
	// CRUD Team member [admin,user]
	// Vote. [user]
	// Receive matches. [user]

	// Restrictions.
	// Only one session at a time for the admin.
	// Only two teams per session: Male and Female.
	// User can be registered only for one team.
	// User can edit his profile only when registration to the team is open.
	// User can vote only once per session. Though, multiple choices are allowed.

	// Session states.
	// Teams creation.
	// Team members registration.
	// Speed dating.
	// Voting.
	// Match making. [when all players vote, or by admin]

	// Corner cases.
	// Some player never votes. Then the admin can start the match making.
}

type fakeUpdate struct{}

func (fakeUpdate) FromAdmin() bool {
	return false
}

func (fakeUpdate) SendMessage(string) {
	return
}
