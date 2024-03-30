package player

type Player struct {
	Account string
	Name    string
	ID      uint8
}

func New(account string, name string, id uint8) *Player {
	return &Player{
		Account: account,
		Name:    name,
		ID:      id,
	}
}

func (p *Player) Choose(id uint8) error {
	return nil
}
