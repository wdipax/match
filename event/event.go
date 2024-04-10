// event converts an incomming notification from messenger to the inner domain object.
package event

type Event struct{}

func New() *Event {
	return &Event{}
}

type Type int

const (
	Unknown Type = iota
	Help
	NewSession
	StartMaleRegistration
	EndMaleRegistration
	StartFemaleRegistration
	EndFemaleRegistration
	AddTeamMember
	TeamMemberName
	TeamMemberNumber
	StartVoting
	Vote
	EndSession
)

// func (e *Event) Command() Type {

// }

// func (e *Event) UserID() string {

// }

// func (e *Event) Payload() string {

// }
