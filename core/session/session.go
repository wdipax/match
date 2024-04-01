package session

import (
	"fmt"

	"github.com/wdipax/match/core/player"
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

func (s *Session) PlayerMatches(id uint8) []uint8 {
	var matches []uint8

	for _, c1 := range s.playerChoices(id) {
		for _, c2 := range s.playerChoices(c1) {
			if c2 == id {
				matches = append(matches, c1)
			}
		}
	}

	return matches
}

func (s *Session) HasPlayer(p *player.Player) bool {
	for _, tm := range s.teams {
		if tm.HasPlayer(p) {
			return true
		}
	}

	return false
}

func (s *Session) playerTeam(id uint8) *team.Team {
	for _, t := range s.teams {
		// TODO: can id collide in different teams?
		if t.HasPlayerWithID(id) {
			return t
		}
	}

	return nil
}

func (s *Session) playerChoices(id uint8) []uint8 {
	var choices []uint8

	for _, v := range s.choices {
		if v[0] == id {
			choices = append(choices, v[1])
		}
	}

	return choices
}
