package state

type State struct{}

func New() *State {
	return &State{}
}

type Update interface {
	FromAdmin() bool
	SendMessage(string)
}

func (s *State) Process(u Update) {
	if u.FromAdmin() {
		u.SendMessage("hello admin")
	} else {
		u.SendMessage("hello user")
	}
}
