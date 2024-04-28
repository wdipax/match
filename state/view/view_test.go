package view_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/state/group/member"
	"github.com/wdipax/match/state/view"
)

func TestPoll(t *testing.T) {
	t.Parallel()

	members := []*member.Member{
		{
			Number: 1,
			Name:   "First",
		},
		{
			Number: 2,
			Name:   "Second",
		},
	}

	data := view.FormPoll(members)

	var p view.Poll

	p.Decode(data)

	s := fmt.Sprint(p.Options)

	assert.Contains(t, s, "1 First")
	assert.Contains(t, s, "2 Second")
}

func TestUserNumberFrom(t *testing.T) {
	t.Parallel()

	n := view.UserNumberFrom("1 First Name")

	assert.Equal(t, 1, n)
}

func TestPollResult(t *testing.T) {
	t.Parallel()

	res := []int{1, 2}

	var r view.PollResult

	data := r.Encode(res)

	r.Decode(data)

	assert.ElementsMatch(t, res, r.Choosen)
}

func TestMatches(t *testing.T) {
	t.Parallel()

	t.Run("it ptints user info", func(t *testing.T) {
		t.Parallel()

		m := []*member.Member{
			{
				Number:  1,
				Name:    "First User",
				Contact: "first",
			},
			{
				Number:  2,
				Name:    "Second User",
				Contact: "second",
			},
		}

		v := view.Matches(m)

		assert.Contains(t, v, strconv.Itoa(m[0].Number))
		assert.Contains(t, v, m[0].Name)
		assert.Contains(t, v, "@"+m[0].Contact)

		assert.Contains(t, v, strconv.Itoa(m[1].Number))
		assert.Contains(t, v, m[1].Name)
		assert.Contains(t, v, "@"+m[1].Contact)
	})

	t.Run("it is empty if no matches", func(t *testing.T) {
		t.Parallel()

		assert.Empty(t, view.Matches(nil))
	})
}
