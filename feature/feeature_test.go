package feature_test

import (
	"strconv"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/adapter"
	"github.com/wdipax/match/event"
	"github.com/wdipax/match/state"
)

const (
	admin = iota + 1
)

func TestSession(t *testing.T) {
	t.Parallel()

	s := state.New(state.StateSettings{
		IsAdmin: func(userID string) bool { return userID == strconv.Itoa(admin) },
	})

	var m fakeMessenger

	a := adapter.New(&m, s)

	a.Process(event.From(&tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				ID: admin,
			},
			Text: "test",
		},
	}))

	assert.True(t, m.received)
}

type fakeMessenger struct {
	received bool
}

func (m *fakeMessenger) Send(userID string, msg string) {
	m.received = true
}
