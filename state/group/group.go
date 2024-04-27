package group

import (
	"github.com/google/uuid"
	"github.com/wdipax/match/state/group/member"
)

type Group struct {
	ID string

	nextNumber func() int
	members    []*member.Member
}

func New() *Group {
	nextNumber := func() func() int {
		var n int
		return func() int {
			n++
			return n
		}
	}()

	return &Group{
		ID:         uuid.NewString(),
		nextNumber: nextNumber,
	}
}

func (g *Group) Add(user int64) int {
	n := g.nextNumber()

	g.members = append(g.members, &member.Member{
		User:   user,
		Number: n,
	})

	return n
}

func (g *Group) HasMember(user int64) bool {
	return g.getMember(user) != nil
}

func (g *Group) SetName(user int64, name string) {
	g.getMember(user).Name = name
}

func (g *Group) Members() []*member.Member {
	return g.members
}

func (g *Group) getMember(user int64) *member.Member {
	for _, m := range g.members {
		if m.User == user {
			return m
		}
	}

	return nil
}
