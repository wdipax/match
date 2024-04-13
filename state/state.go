// state maps incoming requests to the inner logic of the application.
package state

type Core interface {
	NewSession() string
	NewTeam() string
}

type IsAdmin func(userID string) bool

type State struct {
	isAdmin IsAdmin
	core    Core

	userSession map[string]string
}

func New(isAdmin IsAdmin, core Core) *State {
	return &State{
		isAdmin: isAdmin,
		core:    core,

		userSession: make(map[string]string),
	}
}

type Response struct {
	UserID string
	MSG    string
}

func (s *State) Input(userID string, payload string) []*Response {
	return []*Response{
		{
			UserID: userID,
			MSG:    "TODO",
		},
	}
}

func (s *State) Help(userID string) []*Response {
	return nil
}

func (s *State) NewSession(userID string) []*Response {
	if s.isAdmin(userID) {
		s.core.NewSession()
	}

	return nil
}

func (s *State) StartMaleRegistration(userID string) []*Response {
	if !s.isAdmin(userID) {
		return nil
	}

	teamID := s.core.NewTeam()

	return []*Response{
		{
			UserID: userID,
			MSG:    teamID,
		},
	}
}

func (s *State) EndMaleRegistration(userID string) []*Response {
	return nil
}

func (s *State) StartFemaleRegistration(userID string) []*Response {
	if !s.isAdmin(userID) {
		return nil
	}

	teamID := s.core.NewTeam()

	return []*Response{
		{
			UserID: userID,
			MSG:    teamID,
		},
	}
}

func (s *State) EndFemaleRegistration(userID string) []*Response {
	return nil
}

func (s *State) ChangeUserName(userID string) []*Response {
	return nil
}

func (s *State) ChangeUserNumber(userID string) []*Response {
	return nil
}

func (s *State) StartVoting(userID string) []*Response {
	return nil
}

func (s *State) EndSession(userID string) []*Response {
	return nil
}
