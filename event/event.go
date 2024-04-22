package event

type Event struct {
	ChatID              int64
	FromAdmin           bool
	EndTeamRegistration bool
}

func New() *Event {
	return &Event{}
}
