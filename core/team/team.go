package team

import (
	"fmt"

	"github.com/wdipax/match/core/player"
)

type Team struct {
	name  string
	users []*player.Player
}

func New(name string) *Team {
	return &Team{
		name: name,
	}
}

func (t *Team) AddUser(user *player.Player) error {
	for _, v := range t.users {
		if v.Account == user.Account {
			return fmt.Errorf("user already exists: %s", user.Account)
		}
	}

	t.users = append(t.users, user)

	return nil
}
