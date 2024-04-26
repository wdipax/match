package group

import "github.com/google/uuid"

type Group struct {
	ID string
}

func New() *Group {
	return &Group{
		ID: uuid.NewString(),
	}
}
