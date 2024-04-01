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

func (s *Session) Choose(p *player.Player, pID uint8) error {
	t1 := s.playerTeam(p)

	if t1 == nil {
		return fmt.Errorf("player: %d has no team", p.ID)
	}

	teams := s.otherTeams(t1)

	if len(teams) == 0 {
		return fmt.Errorf("player: %d has no team", pID)
	}

	for _, v := range teams {
		if v.HasPlayerWithID(pID) {
			s.choices = append(s.choices, []uint8{p.ID, pID})
		}
	}

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

func (s *Session) playerTeam(p *player.Player) *team.Team {
	for _, t := range s.teams {
		// TODO: can id collide among different teams?
		if t.HasPlayer(p) {
			return t
		}
	}

	return nil
}

func (s *Session) otherTeams(t *team.Team) []*team.Team {
	var res []*team.Team

	for _, v := range s.teams {
		if v.Name != t.Name {
			res = append(res, v)
		}
	}

	return res
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
