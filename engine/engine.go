// engine is a bridge between the outer telegram hendler and the inner core of our application.
package engine

type Engine struct {
	telegram TelegramHandler
	state    StateHandler
}

type TelegramHandler interface {
	Send(msg string)
}

type StateHandler interface {
	Help(userID string) string
	NewSession(userID string) string
	StartMaleRegistration(userID string) string
}

func New(telegram TelegramHandler, state StateHandler) *Engine {
	return &Engine{
		telegram: telegram,
		state:    state,
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
	StartMaleRegistration
)

// Event represents an event from the telegram.
type Event interface {
	Command() Action
	UserID() string
}

func (e *Engine) Process(evt Event) {
	switch evt.Command() {
	case Help:
		e.telegram.Send(e.state.Help(evt.UserID()))
	case NewSession:
		e.telegram.Send(e.state.NewSession(evt.UserID()))
	case StartMaleRegistration:
		e.telegram.Send(e.state.StartMaleRegistration(evt.UserID()))
	}
}
