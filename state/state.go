package state

type State struct{}

func New() *State {
	return nil
}

type Event interface{}

func (s *State) Process(e Event) {}
