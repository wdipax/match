package state

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
	s.users = append(s.users, &User{
		Account: account,
		Name:    name,
		ID:      id,
	})

	return nil
}
