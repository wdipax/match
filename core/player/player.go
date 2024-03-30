package player

type Player struct {
	Account string
	Name    string
	ID      uint8

	session Session
}

type Session interface {
	Choose(p1, p2 uint8) error
}

func New(account string, name string, id uint8, session Session) *Player {
	return &Player{
		Account: account,
		Name:    name,
		ID:      id,

		session: session,
	}
}

func (p *Player) Choose(id uint8) error {
	return p.session.Choose(p.ID, id)
}
