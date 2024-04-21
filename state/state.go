package state

import (
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/response"
)

type State struct{}

func New() *State {
	return &State{}
}

func Process(e *event.Event) *response.Response {
	return nil
}
