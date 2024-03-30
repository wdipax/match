package team

import (
	"fmt"

	"github.com/wdipax/match/core/user"
)

type Team struct {
	name  string
	users []*user.User
}

func New(name string) *Team {
	return &Team{
		name: name,
	}
}

func (t *Team) AddUser(account string, name string, id uint8) error {
	for _, user := range t.users {
		if user.Account == account {
			return fmt.Errorf("user already exists: %s", account)
		}
	}

	t.users = append(t.users, &user.User{
		Account: account,
		Name:    name,
		ID:      id,
	})

	return nil
}
