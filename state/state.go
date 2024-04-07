package state

import (
	"fmt"
	"slices"
)

type State struct {
	admins         []string
	sessionHandler SessionHandler
}

type SessionHandler interface {
	New() string
	Delete(id string)
}

type Settings struct {
	Admins         []string
	SessionHandler SessionHandler
}

func New(settings *Settings) *State {
	return &State{
		admins:         settings.Admins,
		sessionHandler: settings.SessionHandler,
	}
}

func (s *State) NewSession(userID string) error {
	if !slices.Contains(s.admins, userID) {
		return fmt.Errorf("user is not an admin")
	}

	s.sessionHandler.New()

	return nil
}

func (s *State) EndSession(id string) {
	s.sessionHandler.Delete(id)
}
