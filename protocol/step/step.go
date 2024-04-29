package step

const (
	Unknown = iota
	Initialization
	Registration
	KnowEachother
	Voting
	End
)

func Name(s int) string {
	switch s {
	case Initialization:
		return "initialization"
	case Registration:
		return "registration"
	case KnowEachother:
		return "know each other"
	case Voting:
		return "voting"
	case End:
		return "end"
	default:
		return ""
	}
}
