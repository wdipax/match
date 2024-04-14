// state maps incoming requests to the inner logic of the application.
package state

type Core interface {
	NewSession() string
	NewTeam(name string) string
}

type IsAdmin func(userID string) bool

type ResponseMSG func(optional string) string

type StateSettings struct {
	IsAdmin                IsAdmin
	StartTeamMSG           ResponseMSG
	JoinTeamMSG            ResponseMSG
	AdminCanNotJoinTeamMSG ResponseMSG
	EndTeamMSG             ResponseMSG
	VoteReceivedMSG        ResponseMSG
	AdminCanNotVoteMSG     ResponseMSG
	AdminEndSessionMSG     ResponseMSG
	Core                   Core
	MaleTeamName           string
	FemaleTeamName         string
}

type State struct {
	isAdmin                IsAdmin
	startTeamMSG           ResponseMSG
	joinTeamMSG            ResponseMSG
	adminCanNotJoinTeamMSG ResponseMSG
	endTeamMSG             ResponseMSG
	voteReceivedMSG        ResponseMSG
	adminCanNotVoteMSG     ResponseMSG
	adminEndSessionMSG     ResponseMSG
	core                   Core
	maleTeamName           string
	femaleTeamName         string

	sessions    []*session
	teams       []*team
	adminTeams  map[string][]*team
	userSession map[string]*session
}

func New(s StateSettings) *State {
	return &State{
		isAdmin:                s.IsAdmin,
		startTeamMSG:           s.StartTeamMSG,
		joinTeamMSG:            s.JoinTeamMSG,
		adminCanNotJoinTeamMSG: s.AdminCanNotJoinTeamMSG,
		endTeamMSG:             s.EndTeamMSG,
		voteReceivedMSG:        s.VoteReceivedMSG,
		adminCanNotVoteMSG:     s.AdminCanNotVoteMSG,
		adminEndSessionMSG:     s.AdminEndSessionMSG,
		core:                   s.Core,
		maleTeamName:           s.MaleTeamName,
		femaleTeamName:         s.FemaleTeamName,

		adminTeams:  make(map[string][]*team),
		userSession: make(map[string]*session),
	}
}

type team struct {
	id           string
	sessionID    string
	name         string
	registration bool
}

type sessionPhase int

const (
	teamManagement sessionPhase = iota
	voting
)

type session struct {
	id    string
	phase sessionPhase
}

type Response struct {
	UserID string
	MSG    string
}

func (s *State) Input(userID string, payload string) []*Response {
	ss := s.userSession[userID]

	if s.isAdmin(userID) {
		if ss == nil {
			return nil
		}

		if ss.phase == teamManagement {
			return []*Response{
				{
					UserID: userID,
					MSG:    s.adminCanNotJoinTeamMSG(""),
				},
			}
		}

		if ss.phase == voting {
			return []*Response{
				{
					UserID: userID,
					MSG:    s.adminCanNotVoteMSG(""),
				},
			}
		}

		return nil
	}

	if ss != nil && ss.phase == voting {
		return []*Response{
			{
				UserID: userID,
				MSG:    s.voteReceivedMSG(""),
			},
		}
	}

	t := teamByID(s.teams, payload)
	if t == nil {
		return nil
	}

	if !t.registration {
		return nil
	}

	ss = sessionByID(s.sessions, t.sessionID)
	if ss == nil {
		return nil
	}

	s.userSession[userID] = ss

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
	if !s.isAdmin(userID) {
		return nil
	}

	sessionID := s.core.NewSession()

	ss := &session{
		id:    sessionID,
		phase: teamManagement,
	}

	s.sessions = append(s.sessions, ss)

	s.userSession[userID] = ss

	return nil
}

func (s *State) StartMaleRegistration(userID string) []*Response {
	if !s.isAdmin(userID) {
		return nil
	}

	ss, ok := s.userSession[userID]
	if !ok {
		return nil
	}

	teamID := s.core.NewTeam(s.maleTeamName)

	t := &team{
		id:           teamID,
		sessionID:    ss.id,
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

	ss, ok := s.userSession[userID]
	if !ok {
		return nil
	}

	teamID := s.core.NewTeam(s.femaleTeamName)

	t := &team{
		id:           teamID,
		sessionID:    ss.id,
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
	if !s.isAdmin(userID) {
		return nil
	}

	ss, ok := s.userSession[userID]
	if !ok {
		return nil
	}

	ss.phase = voting

	return []*Response{
		{
			UserID: userID,
		},
	}
}

func (s *State) EndSession(userID string) []*Response {
	return []*Response{
		{
			UserID: userID,
			MSG:    s.adminEndSessionMSG(""),
		},
	}
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

func sessionByID(sessions []*session, id string) *session {
	var ss *session

	for _, ss = range sessions {
		if ss.id == id {
			return ss
		}
	}

	return nil
}
