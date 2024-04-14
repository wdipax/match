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

	teams      []*team
	adminTeams map[string][]*team
}

func New(s StateSettings) *State {
	return &State{
		isAdmin:        s.IsAdmin,
		joinTeamMSG:    s.JoinTeamMSG,
		core:           s.Core,
		maleTeamName:   s.MaleTeamName,
		femaleTeamName: s.FemaleTeamName,

		adminTeams: make(map[string][]*team),
	}
}

type team struct {
	id           string
	name         string
	registration bool
}

type Response struct {
	UserID string
	MSG    string
}

func (s *State) Input(userID string, payload string) []*Response {
	if s.isAdmin(userID) {
		return nil
	}

	t := teamByID(s.teams, payload)
	if t == nil {
		return nil
	}

	if !t.registration {
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

	t := &team{
		id:           teamID,
		name:         s.maleTeamName,
		registration: true,
	}

	s.teams = append(s.teams, t)

	s.adminTeams[userID] = append(s.adminTeams[userID], t)

	return []*Response{
		{
			UserID: userID,
			MSG:    teamID,
		},
	}
}

func (s *State) EndMaleRegistration(userID string) []*Response {
	teams, ok := s.adminTeams[userID]
	if !ok {
		return nil
	}

	t := teamByName(teams, s.maleTeamName)
	if t == nil {
		return nil
	}

	t.registration = false

	return []*Response{
		{
			UserID: userID,
			MSG:    t.name,
		},
	}
}

func (s *State) StartFemaleRegistration(userID string) []*Response {
	if !s.isAdmin(userID) {
		return nil
	}

	teamID := s.core.NewTeam(s.femaleTeamName)

	t := &team{
		id:           teamID,
		name:         s.femaleTeamName,
		registration: true,
	}

	s.teams = append(s.teams, t)

	s.adminTeams[userID] = append(s.adminTeams[userID], t)

	return []*Response{
		{
			UserID: userID,
			MSG:    teamID,
		},
	}
}

func (s *State) EndFemaleRegistration(userID string) []*Response {
	teams, ok := s.adminTeams[userID]
	if !ok {
		return nil
	}

	t := teamByName(teams, s.femaleTeamName)
	if t == nil {
		return nil
	}

	t.registration = false

	return []*Response{
		{
			UserID: userID,
			MSG:    t.name,
		},
	}
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

func teamByID(teams []*team, id string) *team {
	var t *team

	for _, t = range teams {
		if t.id == id {
			return t
		}
	}

	return nil
}

func teamByName(teams []*team, name string) *team {
	var t *team

	for _, t = range teams {
		if t.name == name {
			return t
		}
	}

	return nil
}
