package state

import (
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/response"
)

const (
	waitForAdmin = iota
)

type State struct {
	stage int
}

func New() *State {
	return &State{
		stage: waitForAdmin,
	}
}

func Process(e *event.Event) *response.Response {
	return nil
}
