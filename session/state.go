package session

import "fmt"

// TODO: does it belong here?
type User struct {
	Account string
	Name    string
	ID      uint8
	Team    string
}

type Session struct {
	teams []string
	users []*User
}

func New() *Session {
	return &Session{}
}

func (s *Session) NewTeam(name string) error {
	for _, v := range s.teams {
		if v == name {
			return fmt.Errorf("team already exists: %s", name)
		}
	}

	s.teams = append(s.teams, name)

	return nil
}

func (s *Session) Register(account string, name string, id uint8) error {
	for _, v := range s.users {
		if v.Account == account {
			return fmt.Errorf("user already exists: %s", account)
		}
	}

	s.users = append(s.users, &User{
		Account: account,
		Name:    name,
		ID:      id,
	})

	return nil
}

func (s *Session) Choose(ids uint8) error {
	return nil
}
