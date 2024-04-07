package state

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/wdipax/match/core/session"
	"github.com/wdipax/match/core/team"
)

// State manages relationships between objects.
type State struct {
	admins         []string
	sessionHandler SessionHandler

	adminSession             map[string]string
	sessions                 map[string]*session.Session
	sessionTeams             map[string]*mfTeams
	registrationTokenForTeam map[string]*team.Team
}

type mfTeams struct {
	males   *team.Team
	females *team.Team
}

// SessionHandler interacts with core session.
// TODO: make it responsible for interacting wit core session.
type SessionHandler interface {
	New() string
	Delete(id string)
	Help(admin bool) string
	StartRegistration(teamID string) (string, error)
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
		sessions:     make(map[string]*session.Session),
	}
}

func (s *State) NewSession(userID string) error {
	if !s.isAdmin(userID) {
		return fmt.Errorf("user is not an admin")
	}

	if _, exists := s.adminSession[userID]; exists {
		return fmt.Errorf("a session already exists")
	}

	sessionID := s.sessionHandler.New()

	s.adminSession[userID] = sessionID

	se := session.New()

	s.sessions[sessionID] = se

	males := team.New("males", se)

	err := se.AddTeam(males)
	if err != nil {
		return fmt.Errorf("creating males team: %w", err)
	}

	females := team.New("females", se)

	err = se.AddTeam(females)
	if err != nil {
		return fmt.Errorf("creating females team: %w", err)
	}

	s.sessionTeams[sessionID] = &mfTeams{
		males:   males,
		females: females,
	}

	return nil
}

func (s *State) EndSession(id string) {
	s.sessionHandler.Delete(id)
}

func (s *State) Help(userID string) string {
	return s.sessionHandler.Help(s.isAdmin(userID))
}

func (s *State) StartMaleRegistration(userID string) (string, error) {
	if !s.isAdmin(userID) {
		return "", fmt.Errorf("user is not an admin")
	}

	sessionID, exists := s.adminSession[userID]
	if !exists {
		return "", fmt.Errorf("admin does not have an active session")
	}

	teams, exists := s.sessionTeams[sessionID]
	if !exists {
		return "", fmt.Errorf("session does not have teams")
	}

	token := uuid.NewString()

	s.registrationTokenForTeam[token] = teams.males

	return token, nil
}

func (s *State) isAdmin(userID string) bool {
	return slices.Contains(s.admins, userID)
}
