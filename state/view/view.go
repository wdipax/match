package view

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/wdipax/match/protocol/step"
	"github.com/wdipax/match/state/group/member"
)

func GroupMembers(members []*member.Member, stage int) string {
	switch stage {
	case step.Registration:
		sortMembers(members)

		var acc []string

		for _, m := range members {
			acc = append(acc, fmt.Sprintf("%d\t%s", m.Number, m.Name))
		}

		return strings.Join(acc, "\n")
	default:
		return ""
	}
}

func sortMembers(members []*member.Member) {
	slices.SortFunc(members, func(a, b *member.Member) int {
		if n := cmp.Compare(a.Number, b.Number); n != 0 {
			return n
		}

		return cmp.Compare(a.Name, b.Name)
	})
}
