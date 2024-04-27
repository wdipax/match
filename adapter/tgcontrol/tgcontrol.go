package tgcontrol

import "github.com/wdipax/match/protocol/step"

func Stat(stage int) string {
	switch stage {
	case step.Registration:
		return "участники"
	default:
		return ""
	}
}

func Next(stage int) string {
	switch stage {
	case step.Registration:
		return "начать знакомство"
	case step.KnowEachother:
		return "начать голосование"
	default:
		return ""
	}
}

func Back(stage int) string {
	switch stage {
	case step.KnowEachother:
		return "вернуться к регистрации"
	default:
		return ""
	}
}
