package session

import (
	"fmt"

	"github.com/wdipax/match/core/team"
)

type Session struct {
	teams   []*team.Team
	choices [][]uint8
}

func New() *Session {
	return &Session{}
}

func (s *Session) AddTeam(t *team.Team) error {
	for _, v := range s.teams {
		if v.Name == t.Name {
			return fmt.Errorf("team already exists: %s", t.Name)
		}
	}

	s.teams = append(s.teams, t)

	return nil
}

func (s *Session) Choose(p1, p2 uint8) error {
	t1 := s.playerTeam(p1)

	if t1 == nil {
		return fmt.Errorf("player: %d has no team", p1)
	}

	t2 := s.playerTeam(p2)

	if t2 == nil {
		return fmt.Errorf("player: %d has no team", p2)
	}

	if t1.Name == t2.Name {
		return fmt.Errorf("can not chose a player from the same team")
	}

	s.choices = append(s.choices, []uint8{p1, p2})

	return nil
}

func (s *Session) ComputeMatches(p uint8) []uint8 {
	return nil
}

func (s *Session) playerTeam(id uint8) *team.Team {
	for _, t := range s.teams {
		if t.HasPlayer(id) {
			return t
		}
	}

	return nil
}
