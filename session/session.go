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
	users []*User
}

type User struct {
	ID     int64
	name   string
	number int
}

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrNoUser Error = "no such user"
)

func (s *Session) SetUserName(userID int64, name string) error {
	u := s.getUser(userID)
	if u == nil {
		return ErrNoUser
	}

	u.name = name

	return nil
}

func (s *Session) SetUserNumber(userID int64, num int) error {
	u := s.getUser(userID)
	if u == nil {
		return ErrNoUser
	}

	u.number = num

	return nil
}

func (s *Session) GetAllUsers() []*User {
	return append(s.boys.users, s.girls.users...)
}

func (s *Session) getUser(id int64) *User {
	for _, u := range s.boys.users {
		if u.ID == id {
			return u
		}
	}

	for _, u := range s.girls.users {
		if u.ID == id {
			return u
		}
	}

	return nil
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
		if u.ID == userID {
			return false
		}
	}

	t.users = append(t.users, &User{
		ID: userID,
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
