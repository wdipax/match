package user

type User struct {
	Account string
	Name    string
	ID      uint8
}

func New(account string, name string, id uint8) *User {
	return &User{
		Account: account,
		Name:    name,
		ID:      id,
	}
}
