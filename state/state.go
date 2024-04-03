package state

type State struct{}

func New() *State {
	return &State{}
}

type Update interface{}

func (s *State) Process(u Update) {

}
