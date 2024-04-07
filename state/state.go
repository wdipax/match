package state

import (
	"fmt"
	"slices"
)

type State struct {
	admins         []string
	sessionHandler SessionHandler

	adminSession map[string]string
}

type SessionHandler interface {
	New() string
	Delete(id string)
	Help(admin bool) string
}

type Settings struct {
	Admins         []string
	SessionHandler SessionHandler
}

func New(settings *Settings) *State {
	return &State{
		admins:         settings.Admins,
		sessionHandler: settings.SessionHandler,

		adminSession: make(map[string]string),
	}
}

func (s *State) NewSession(userID string) error {
	if !s.isAdmin(userID) {
		return fmt.Errorf("user is not an admin")
	}

	if _, exists := s.adminSession[userID]; exists {
		return fmt.Errorf("a session already exists")
	}

	s.adminSession[userID] = s.sessionHandler.New()

	return nil
}

func (s *State) EndSession(id string) {
	s.sessionHandler.Delete(id)
}

func (s *State) Help(userID string) string {
	return s.sessionHandler.Help(s.isAdmin(userID))
}

func (s *State) isAdmin(userID string) bool {
	return slices.Contains(s.admins, userID)
}
