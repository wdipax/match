package state

import "fmt"

// TODO: does it belong here?
type User struct {
	Account string
	Name    string
	ID      uint8
}

type State struct {
	users []*User
}

func New() *State {
	return &State{}
}

func (s *State) Register(account string, name string, id uint8) error {
	for _, v := range s.users {
		if v.Account == account {
			return fmt.Errorf("user with that account already exists: %s", account)
		}
	}
	
	s.users = append(s.users, &User{
		Account: account,
		Name:    name,
		ID:      id,
	})

	return nil
}
