// engine is a bridge between the outer telegram hendler and the inner core of our application.
package engine

type Engine struct {
	state StateHandler
}

type StateHandler interface {
	NewSession(userID string) error
}

func New(state StateHandler) *Engine {
	return &Engine{
		state: state,
	}
}

// Telegram represents a telegram handler.
type Telegram interface{}

// Action represenst an event action.
type Action int

const (
	Unknown    Action = 0
	NewSession Action = 1
)

// Event represents an event from the telegram.
type Event interface {
	Command() Action
	UserID() string
}

func (s *Engine) Process(e Event) {
	switch e.Command() {
	case NewSession:
		s.state.NewSession(e.UserID())
	}
}
