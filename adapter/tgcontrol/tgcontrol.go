package tgcontrol

import "github.com/wdipax/match/protocol/step"

func Stat(stage int) string {
	switch stage {
	case step.Registration:
		return members
	case step.Voting:
		return statistics
	default:
		return ""
	}
}

func Next(stage int) string {
	switch stage {
	case step.Registration:
		return startDating
	case step.KnowEachother:
		return startVoting
	case step.Voting:
		return endVoting
	default:
		return ""
	}
}

func Back(stage int) string {
	switch stage {
	case step.KnowEachother:
		return backToRegistration
	default:
		return ""
	}
}

func Repeat(stage int) string {
	switch stage {
	case step.Voting:
		return revote
	default:
		return ""
	}
}
