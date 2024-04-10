// state maps incoming requests to the inner logic of the application.
package state

type State struct {
	admins []string
	core   CoreHandler

	userSession map[string]string
}

type CoreHandler interface{}

type Settings struct {
	Admins []string
	Core   CoreHandler
}

func New(settings *Settings) *State {
	return &State{
		admins: settings.Admins,
		core:   settings.Core,

		userSession: make(map[string]string),
	}
}

type Response struct {
	UserID string
	MSG    string
}

func (s *Settings) Help(userID string) []*Response {
	return nil
}

func (s *Settings) NewSession(userID string) []*Response {
	return nil
}

func (s *Settings) StartMaleRegistration(userID string) []*Response {
	return nil
}

func (s *Settings) EndMaleRegistration(userID string) []*Response {
	return nil
}

func (s *Settings) StartFemaleRegistration(userID string) []*Response {
	return nil
}

func (s *Settings) EndFemaleRegistration(userID string) []*Response {
	return nil
}

func (s *Settings) AddTeamMember(userID string, teamID string) []*Response {
	return nil
}

func (s *Settings) TeamMemberName(userID string, name string) []*Response {
	return nil
}

func (s *Settings) TeamMemberNumber(userID string, number string) []*Response {
	return nil
}

func (s *Settings) StartVoting(userID string) []*Response {
	return nil
}

func (s *Settings) Vote(userID string, poll string) []*Response {
	return nil
}

func (s *Settings) EndSession(userID string) []*Response {
	return nil
}
