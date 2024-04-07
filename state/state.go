package state

type State struct {
	sessionHandler SessionHandler
}

type SessionHandler interface {
	New()
}

func New(sessionHandler SessionHandler) *State {
	return &State{
		sessionHandler: sessionHandler,
	}
}

func (s *State) NewSession() {
	s.sessionHandler.New()
}
