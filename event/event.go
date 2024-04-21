package event

type Event struct{}

func New() *Event {
	return &Event{}
}

func (e *Event) FromAdmin() bool {
	return true
}
