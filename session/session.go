package session

import "github.com/google/uuid"

type Session struct {
	boys  *team
	girls *team
}

func New() *Session {
	return &Session{}
}

type team struct {
	id    string
	users []*user
}

type user struct {
	id int64
}

func (s *Session) JoinTeam(teamID string, userID int64) bool {
	switch {
	case s.boys.id == teamID:
		return joinTeam(s.boys, userID)
	case s.girls.id == teamID:
		return joinTeam(s.girls, userID)
	default:
		return false
	}
}

func joinTeam(t *team, userID int64) bool {
	for _, u := range t.users {
		if u.id == userID {
			return false
		}
	}

	t.users = append(t.users, &user{
		id: userID,
	})

	return true
}

func (s *Session) CreateBoysTeam() string {
	t := team{
		id: uuid.NewString(),
	}

	s.boys = &t

	return t.id
}

func (s *Session) CreateGirlsTeam() string {
	t := team{
		id: uuid.NewString(),
	}

	s.girls = &t

	return t.id
}
