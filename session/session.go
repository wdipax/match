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
	id string
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
