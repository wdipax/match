package state

import "github.com/wdipax/match/event"

type State struct{}

func New() *State {
	return &State{}
}

func Process(e *event.Event) {

}
