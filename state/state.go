package state

type State struct {
	sessionHandler SessionHandler
}

type SessionHandler interface {
	New() string
	Delete(id string)
}

func New(sessionHandler SessionHandler) *State {
	return &State{
		sessionHandler: sessionHandler,
	}
}

func (s *State) NewSession() string {
	return s.sessionHandler.New()
}

func (s *State) EndSession(id string) {
	s.sessionHandler.Delete(id)
}
