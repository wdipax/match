package state

import (
	"fmt"
	"slices"
)

// State manages relationships between objects.
type State struct {
	admins         []string
	sessionHandler SessionHandler

	userSession map[string]string
}

// SessionHandler interacts with the application core.
type SessionHandler interface {
	New() string
	Delete(sessionID string)
	Help(admin bool) string
	StartMaleRegistration(sessionID string) (string, error)
}

type Settings struct {
	Admins         []string
	SessionHandler SessionHandler
}

func New(settings *Settings) *State {
	return &State{
		admins:         settings.Admins,
		sessionHandler: settings.SessionHandler,

		userSession: make(map[string]string),
	}
}

func (s *State) NewSession(userID string) error {
	if !s.isAdmin(userID) {
		return fmt.Errorf("user is not an admin")
	}

	if _, exists := s.userSession[userID]; exists {
		return fmt.Errorf("session already exists")
	}

	s.userSession[userID] = s.sessionHandler.New()

	return nil
}

func (s *State) EndSession(sessionID string) {
	s.sessionHandler.Delete(sessionID)
}

func (s *State) Help(userID string) string {
	return s.sessionHandler.Help(s.isAdmin(userID))
}

func (s *State) StartMaleRegistration(userID string) (string, error) {
	if !s.isAdmin(userID) {
		return "", fmt.Errorf("user is not an admin")
	}

	sessionID, exists := s.userSession[userID]
	if !exists {
		return "", fmt.Errorf("user does not have an active session")
	}

	teamID, err := s.sessionHandler.StartMaleRegistration(sessionID)
	if err != nil {
		return "", fmt.Errorf("starting male registration: %w", err)
	}

	return teamID, nil
}

func (s *State) isAdmin(userID string) bool {
	return slices.Contains(s.admins, userID)
}
