// state maps incoming requests to the inner logic of the application.
package state

type Core interface {
	NewSession() string
	NewTeam(name string) string
	Poll(teamID string) string
	Vote(userID string, voice string)
}

type IsAdmin func(userID string) bool

type ResponseMSG func(optional string) string

type StateSettings struct {
	IsAdmin         IsAdmin
	StartSessionMSG ResponseMSG
	// StartTeamMSG           ResponseMSG
	JoinTeamMSG            ResponseMSG
	AdminCanNotJoinTeamMSG ResponseMSG
	EndTeamMSG             ResponseMSG
	StartVotingMSG         ResponseMSG
	VoteReceivedMSG        ResponseMSG
	AdminCanNotVoteMSG     ResponseMSG
	EndSessionMSG          ResponseMSG
	Core                   Core
	MaleTeamName           string
	FemaleTeamName         string
}

type State struct {
	isAdmin         IsAdmin
	startSessionMSG ResponseMSG
	// startTeamMSG           ResponseMSG
	joinTeamMSG            ResponseMSG
	adminCanNotJoinTeamMSG ResponseMSG
	endTeamMSG             ResponseMSG
	startVotingMSG         ResponseMSG
	voteReceivedMSG        ResponseMSG
	adminCanNotVoteMSG     ResponseMSG
	endSessionMSG          ResponseMSG
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
		isAdmin:         s.IsAdmin,
		startSessionMSG: s.StartSessionMSG,
		// startTeamMSG:           s.StartTeamMSG,
		joinTeamMSG:            s.JoinTeamMSG,
		adminCanNotJoinTeamMSG: s.AdminCanNotJoinTeamMSG,
		endTeamMSG:             s.EndTeamMSG,
		voteReceivedMSG:        s.VoteReceivedMSG,
		adminCanNotVoteMSG:     s.AdminCanNotVoteMSG,
		startVotingMSG:         s.StartVotingMSG,
		endSessionMSG:          s.EndSessionMSG,
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
	users        []*user
}

type user struct {
	id    string
	voted bool
}

func (t *team) hasUser(id string) bool {
	return t.user(id) != nil
}

func (t *team) user(id string) *user {
	for _, user := range t.users {
		if user.id == id {
			return user
		}
	}

	return nil
}

func (t *team) addUser(id string) {
	t.users = append(t.users, &user{
		id: id,
	})
}

func vote(tms []*team, userID string) {
	var u *user

	for _, t := range tms {
		if u = t.user(userID); u != nil {
			u.voted = true

			return
		}
	}
}

func allUsersVoted(tms []*team) bool {
	for _, t := range tms {
		for _, u := range t.users {
			if !u.voted {
				return false
			}
		}
	}

	return true
}

type sessionPhase int

const (
	teamManagement sessionPhase = iota
	voting
)

type session struct {
	id      string
	phase   sessionPhase
	adminID string
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
		tms := s.adminTeams[ss.adminID]

		vote(tms, userID)

		res := []*Response{
			{
				UserID: userID,
				MSG:    s.voteReceivedMSG(""),
			},
		}

		if allUsersVoted(tms) {
			// defer deleteSession(ss.id)

			for _, t := range tms {
				for _, u := range t.users {
					res = append(res, &Response{
						UserID: u.id,
						MSG:    "here is your matches",
					})
				}
			}

			res = append(res, &Response{
				UserID: ss.adminID,
				MSG:    s.endSessionMSG(""),
			})
		}

		return res
	}

	t := teamByID(s.teams, payload)
	if t == nil {
		return nil
	}

	if !t.registration {
		return nil
	}

	if t.hasUser(userID) {
		return nil
	}

	ss = sessionByID(s.sessions, t.sessionID)
	if ss == nil {
		return nil
	}

	s.userSession[userID] = ss

	t.addUser(userID)

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
		id:      sessionID,
		phase:   teamManagement,
		adminID: userID,
	}

	s.sessions = append(s.sessions, ss)

	s.userSession[userID] = ss

	return []*Response{
		{
			UserID: userID,
			MSG:    s.startSessionMSG(""),
		},
	}
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
			MSG:    s.endTeamMSG(t.name),
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
			MSG:    s.endTeamMSG(t.name),
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

	tms := s.adminTeams[userID]

	res := []*Response{
		{
			UserID: userID,
			MSG:    s.startVotingMSG(""),
		},
	}

	type teamPole struct {
		teamID string
		pole   string
	}

	var polls []teamPole

	for _, t := range tms {
		polls = append(polls, teamPole{
			teamID: t.id,
			pole:   s.core.Poll(t.id),
		})
	}

	for _, t := range tms {
		for _, u := range t.users {
			for _, p := range polls {
				if p.teamID != t.id {
					res = append(res, &Response{
						UserID: u.id,
						MSG:    p.pole,
					})
				}
			}
		}
	}

	return res
}

func (s *State) EndSession(userID string) []*Response {
	if !s.isAdmin(userID) {
		return nil
	}

	return []*Response{
		{
			UserID: userID,
			MSG:    s.endSessionMSG(""),
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
