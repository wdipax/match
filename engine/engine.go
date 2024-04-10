// engine is a bridge between the outer telegram hendler and the inner core of our application.
package engine

type Engine struct {
	telegram TelegramHandler
	state    StateHandler
}

type TelegramHandler interface {
	Send(userID string, msg string)
}

type Message struct {
	UserID string
	MSG    string
}

type StateHandler interface {
	Help(userID string) []*Message
	NewSession(userID string) []*Message
	StartMaleRegistration(userID string) []*Message
	EndMaleRegistration(userID string) []*Message
	StartFemaleRegistration(userID string) []*Message
	EndFemaleRegistration(userID string) []*Message
	AddTeamMember(userID string, teamID string) []*Message
	TeamMemberName(userID string, name string) []*Message
	TeamMemberNumber(userID string, number string) []*Message
	StartVoting(userID string) []*Message
	Vote(userID string, poll string) []*Message
	EndSession(userID string) []*Message
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
	EndSession
	StartMaleRegistration
	EndMaleRegistration
	StartFemaleRegistration
	EndFemaleRegistration
	AddTeamMember
	TeamMemberName
	TeamMemberNumber
	StartVoting
	Vote
)

// Event represents an event from the telegram.
type Event interface {
	Command() Action
	UserID() string
	Payload() string
}

func (e *Engine) Process(evt Event) {
	userID := evt.UserID()
	payload := evt.Payload()

	var responses []*Message

	switch evt.Command() {
	case Help:
		responses = e.state.Help(userID)
	case NewSession:
		responses = e.state.NewSession(userID)
	case StartMaleRegistration:
		responses = e.state.StartMaleRegistration(userID)
	case EndMaleRegistration:
		responses = e.state.EndMaleRegistration(userID)
	case StartFemaleRegistration:
		responses = e.state.StartFemaleRegistration(userID)
	case EndFemaleRegistration:
		responses = e.state.EndFemaleRegistration(userID)
	case AddTeamMember:
		responses = e.state.AddTeamMember(userID, payload)
	case TeamMemberName:
		responses = e.state.TeamMemberName(userID, payload)
	case TeamMemberNumber:
		responses = e.state.TeamMemberNumber(userID, payload)
	case StartVoting:
		responses = e.state.StartVoting(userID)
	case Vote:
		responses = e.state.Vote(userID, payload)
	case EndSession:
		responses = e.state.EndSession(userID)
	}

	for _, re := range responses {
		e.telegram.Send(re.UserID, re.MSG)
	}
}
