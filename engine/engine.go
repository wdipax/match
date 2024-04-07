// engine is a bridge between the outer telegram hendler and the inner core of our application.
package engine

type Engine struct {
	state StateHandler
}

type StateHandler interface {
	NewSession(userID string) error
	Help(userID string) string
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
	Unknown Action = iota
	Help
	NewSession
)

// Event represents an event from the telegram.
type Event interface {
	Command() Action
	UserID() string
}

func (e *Engine) Process(evt Event) {
	switch evt.Command() {
	case Help:
		e.state.Help(evt.UserID())
	case NewSession:
		e.state.NewSession(evt.UserID())
	}
}
