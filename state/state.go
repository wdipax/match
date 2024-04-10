// state maps incoming requests to the inner logic of the application.
package state

type CoreHandler interface{}

type State struct {
	admins []string
	core   CoreHandler

	userSession map[string]string
}

func New(admins []string, core CoreHandler) *State {
	return &State{
		admins: admins,
		core:   core,

		userSession: make(map[string]string),
	}
}

type Response struct {
	UserID string
	MSG    string
}

func (s *State) Help(userID string) []*Response {
	return nil
}

func (s *State) NewSession(userID string) []*Response {
	return nil
}

func (s *State) StartMaleRegistration(userID string) []*Response {
	return nil
}

func (s *State) EndMaleRegistration(userID string) []*Response {
	return nil
}

func (s *State) StartFemaleRegistration(userID string) []*Response {
	return nil
}

func (s *State) EndFemaleRegistration(userID string) []*Response {
	return nil
}

func (s *State) AddTeamMember(userID string, teamID string) []*Response {
	return nil
}

func (s *State) TeamMemberName(userID string, name string) []*Response {
	return nil
}

func (s *State) TeamMemberNumber(userID string, number string) []*Response {
	return nil
}

func (s *State) StartVoting(userID string) []*Response {
	return nil
}

func (s *State) Vote(userID string, poll string) []*Response {
	return nil
}

func (s *State) EndSession(userID string) []*Response {
	return nil
}
