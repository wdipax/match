package team

import (
	"fmt"

	"github.com/wdipax/match/core/player"
)

type Team struct {
	Name string

	players []*player.Player
}

func New(name string) *Team {
	return &Team{
		Name: name,
	}
}

func (t *Team) AddPlayer(p *player.Player) error {
	for _, v := range t.players {
		if v.Account == p.Account {
			return fmt.Errorf("player already exists: %s", p.Account)
		}

		if v.ID == p.ID {
			return fmt.Errorf("player id is already taken: %d", p.ID)
		}
	}

	t.players = append(t.players, p)

	return nil
}

func (t *Team) HasPlayer(id uint8) bool {
	for _, p := range t.players {
		if p.ID == id {
			return true
		}
	}

	return false
}
