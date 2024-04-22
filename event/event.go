package event

type Event struct {
	FromAdmin           bool
	EndTeamRegistration bool
}

func New() *Event {
	return &Event{}
}
