package session

import (
	"fmt"
)

type Session struct {
	teams []string
}

func New() *Session {
	return &Session{}
}

func (s *Session) AddTeam(name string) error {
	for _, v := range s.teams {
		if v == name {
			return fmt.Errorf("team already exists: %s", name)
		}
	}

	s.teams = append(s.teams, name)

	return nil
}
