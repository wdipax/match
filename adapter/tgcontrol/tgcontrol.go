package tgcontrol

import "github.com/wdipax/match/protocol/step"

func Stat(stage int) string {
	switch stage {
	case step.Registration:
		return "участники"
	case step.Voting:
		return "статистика"
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
	case step.Voting:
		return "подвести итоги"
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

func Repeat(stage int) string {
	switch stage {
	case step.Voting:
		return "изменить выбор"
	default:
		return ""
	}
}
