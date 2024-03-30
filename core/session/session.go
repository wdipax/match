package session

import (
	"fmt"

	"github.com/wdipax/match/core/team"
)

type Session struct {
	teams []*team.Team
}

func New() *Session {
	return &Session{}
}

func (s *Session) AddTeam(team *team.Team) error {
	for _, v := range s.teams {
		if v.Name == team.Name {
			return fmt.Errorf("team already exists: %s", team.Name)
		}
	}

	s.teams = append(s.teams, team)

	return nil
}

func (s *Session) Choose(p1, p2 uint8) error {

	return nil
}
