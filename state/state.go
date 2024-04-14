// state maps incoming requests to the inner logic of the application.
package state

type Core interface {
	NewSession() string
	NewTeam(name string) string
}

type IsAdmin func(userID string) bool

type StateSettings struct {
	IsAdmin        IsAdmin
	Core           Core
	MaleTeamName   string
	FemaleTeamName string
}

type State struct {
	isAdmin        IsAdmin
	core           Core
	maleTeamName   string
	femaleTeamName string

	userSession map[string]string
}

func New(s StateSettings) *State {
	return &State{
		isAdmin:        s.IsAdmin,
		core:           s.Core,
		maleTeamName:   s.MaleTeamName,
		femaleTeamName: s.FemaleTeamName,

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
			MSG:    s.maleTeamName,
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

	teamID := s.core.NewTeam("male")

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

	teamID := s.core.NewTeam("male")

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
