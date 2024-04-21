package event

type Event struct{}

func New() *Event {
	return &Event{}
}

func (e *Event) FromAdmin() bool {
	return true
}

func (s *Event) EndTeamRegistration() bool {
	return true
}
