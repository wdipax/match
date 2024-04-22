package event

type Event struct {
	ChatID    int64
	FromAdmin bool
	TeamID    string
}

func New(chatID int64, fromAdmin bool, teamID string) *Event {
	return &Event{
		ChatID:    chatID,
		FromAdmin: fromAdmin,
		TeamID:    teamID,
	}
}
