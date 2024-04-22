package event

type Event struct {
	ChatID              int64
	FromAdmin           bool
	EndTeamRegistration bool
}

func New(chatID int64, fromAdmin bool) *Event {
	return &Event{
		ChatID:    chatID,
		FromAdmin: fromAdmin,
	}
}
