package event

type Event struct {
	ChatID    int64
	Text      string
	FromAdmin bool
	TeamID    string
}

func New(chatID int64, text string, fromAdmin bool, teamID string) *Event {
	return &Event{
		ChatID:    chatID,
		Text:      text,
		FromAdmin: fromAdmin,
		TeamID:    teamID,
	}
}
