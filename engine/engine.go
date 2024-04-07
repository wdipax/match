// engine is a bridge between the outer telegram hendler and the inner core of our application.
package engine

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

// Telegram represents a telegram handler.
type Telegram interface{}

// Event represents an event from the telegram.
type Event interface {
	FromAdmin() bool
	SendMessage(string)
}

func (s *Engine) Process(e Event) {
	if e.FromAdmin() {
		e.SendMessage("hello admin")
	} else {
		e.SendMessage("hello user")
	}
}
