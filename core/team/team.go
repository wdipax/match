package team

import (
	"fmt"

	"github.com/wdipax/match/core/player"
)

type Team struct {
	Name string

	players  []*player.Player
	allTeams Session
}

type Session interface {
	HasPlayer(*player.Player) bool
}

func New(name string, allTeams Session) *Team {
	return &Team{
		Name:     name,
		allTeams: allTeams,
	}
}

func (t *Team) AddPlayer(p *player.Player) error {
	if t.allTeams.HasPlayer(p) {
		return fmt.Errorf("player already exists: %s", p.Account)
	}

	if t.HasPlayerWithID(p.ID) {
		return fmt.Errorf("player id is already taken: %d", p.ID)
	}

	t.players = append(t.players, p)

	return nil
}

func (t *Team) HasPlayer(p *player.Player) bool {
	for _, v := range t.players {
		if v.Account == p.Account {
			return true
		}
	}

	return false
}

func (t *Team) HasPlayerWithID(id uint8) bool {
	for _, p := range t.players {
		if p.ID == id {
			return true
		}
	}

	return false
}
