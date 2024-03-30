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

func (t *Team) AddUser(user *user.User) error {
	for _, v := range t.users {
		if v.Account == user.Account {
			return fmt.Errorf("user already exists: %s", user.Account)
		}
	}

	t.users = append(t.users, user)

	return nil
}
