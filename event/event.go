package event

const (
	Unknown = iota
	PreviousStage
	NextStage
	Statistics
)

type Event struct {
	ChatID    int64
	Text      string
	FromAdmin bool
	TeamID    string
	Command   int
}

func New(chatID int64, text string, fromAdmin bool, teamID string, command int) *Event {
	return &Event{
		ChatID:    chatID,
		Text:      text,
		FromAdmin: fromAdmin,
		TeamID:    teamID,
		Command:   command,
	}
}
