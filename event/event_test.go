package event_test

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/event"
)

func TestEvent_Command(t *testing.T) {
	t.Parallel()

	t.Run("help", func(t *testing.T) {
		t.Parallel()

		u := &tgbotapi.Update{
			Message: &tgbotapi.Message{
				Text: "/help",
			},
		}

		e := event.From(u)

		assert.Equal(t, event.Help, e.Command())
	})

	t.Run("new session", func(t *testing.T) {
		t.Parallel()

		u := &tgbotapi.Update{
			Message: &tgbotapi.Message{
				Text: "/new_session",
			},
		}

		e := event.From(u)

		assert.Equal(t, event.NewSession, e.Command())
	})
}
