// engine is a bridge between the outer telegram hendler and the inner core of our application.
package engine

type Engine struct {
	telegram TelegramHandler
	state    StateHandler
}

type TelegramHandler interface {
	Send(userID string, msg string)
}

type StateHandler interface {
	Help(userID string) string
	NewSession(userID string) string
	StartMaleRegistration(userID string) string
	EndMaleRegistration(userID string) string
	StartFemaleRegistration(userID string) string
	EndFemaleRegistration(userID string) string
	AddTeamMember(userID string, teamID string) string
	TeamMemberName(userID string, name string) string
	TeamMemberNumber(userID string, number string) string
	StartVoting(userID string) string
	Vote(userID string, poll string) string
	EndSession(userID string) string
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

	switch evt.Command() {
	case Help:
		e.telegram.Send(userID, e.state.Help(userID))
	case NewSession:
		e.telegram.Send(userID, e.state.NewSession(userID))
	case StartMaleRegistration:
		e.telegram.Send(userID, e.state.StartMaleRegistration(userID))
	case EndMaleRegistration:
		e.telegram.Send(userID, e.state.EndMaleRegistration(userID))
	case StartFemaleRegistration:
		e.telegram.Send(userID, e.state.StartFemaleRegistration(userID))
	case EndFemaleRegistration:
		e.telegram.Send(userID, e.state.EndFemaleRegistration(userID))
	case AddTeamMember:
		e.telegram.Send(userID, e.state.AddTeamMember(userID, evt.Payload()))
	case TeamMemberName:
		e.telegram.Send(userID, e.state.TeamMemberName(userID, evt.Payload()))
	case TeamMemberNumber:
		e.telegram.Send(userID, e.state.TeamMemberNumber(userID, evt.Payload()))
	case StartVoting:
		e.telegram.Send(userID, e.state.StartVoting(userID))
	case Vote:
		e.telegram.Send(userID, e.state.Vote(userID, evt.Payload()))
	case EndSession:
		// TODO: send voting results.
		e.telegram.Send(userID, e.state.EndSession(userID))
	}
}
