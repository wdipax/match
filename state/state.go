// state maps incoming requests to the inner logic of the application.
package state

type Core interface {
	NewSession() string
	NewTeam(name string) string
}

type IsAdmin func(userID string) bool

type JoinTeamMSG func(teamName string) string

type StateSettings struct {
	IsAdmin        IsAdmin
	JoinTeamMSG    JoinTeamMSG
	Core           Core
	MaleTeamName   string
	FemaleTeamName string
}

type State struct {
	isAdmin        IsAdmin
	joinTeamMSG    JoinTeamMSG
	core           Core
	maleTeamName   string
	femaleTeamName string

	teams []*team
}

func New(s StateSettings) *State {
	return &State{
		isAdmin:        s.IsAdmin,
		joinTeamMSG:    s.JoinTeamMSG,
		core:           s.Core,
		maleTeamName:   s.MaleTeamName,
		femaleTeamName: s.FemaleTeamName,
	}
}

type team struct {
	id   string
	name string
}

type Response struct {
	UserID string
	MSG    string
}

func (s *State) Input(userID string, payload string) []*Response {
	if s.isAdmin(userID) {
		return nil
	}

	var t *team

	for _, v := range s.teams {
		if v.id == payload {
			t = v
		}
	}

	if t == nil {
		return nil
	}

	return []*Response{
		{
			UserID: userID,
			MSG:    s.joinTeamMSG(t.name),
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

	teamID := s.core.NewTeam(s.maleTeamName)

	s.teams = append(s.teams, &team{
		id:   teamID,
		name: s.maleTeamName,
	})

	return []*Response{
		{
			UserID: userID,
			MSG:    teamID,
		},
	}
}

func (s *State) EndMaleRegistration(userID string) []*Response {
	return []*Response{
		{
			UserID: userID,
			MSG:    s.maleTeamName,
		},
	}
}

func (s *State) StartFemaleRegistration(userID string) []*Response {
	if !s.isAdmin(userID) {
		return nil
	}

	teamID := s.core.NewTeam(s.femaleTeamName)

	s.teams = append(s.teams, &team{
		id:   teamID,
		name: s.femaleTeamName,
	})

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
