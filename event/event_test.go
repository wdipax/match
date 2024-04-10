package event_test

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"github.com/wdipax/match/event"
)

func TestEvent_Command(t *testing.T) {
	t.Parallel()

	for _, testCase := range [...]struct {
		operation string
		command   string
		expect    event.Type
	}{
		{
			operation: "input",
			command:   "some text",
			expect:    event.Input,
		},
		{
			operation: "help",
			command:   "/help",
			expect:    event.Help,
		},
		{
			operation: "new session",
			command:   "/new_session",
			expect:    event.NewSession,
		},
		{
			operation: "start male registration",
			command:   "/start_male_registration",
			expect:    event.StartMaleRegistration,
		},
		{
			operation: "end male registration",
			command:   "/end_male_registration",
			expect:    event.EndMaleRegistration,
		},
		{
			operation: "start female registration",
			command:   "/start_female_registration",
			expect:    event.StartFemaleRegistration,
		},
		{
			operation: "end female registration",
			command:   "/end_female_registration",
			expect:    event.EndFemaleRegistration,
		},
		{
			operation: "change user name",
			command:   "/change_user_name",
			expect:    event.ChangeUserName,
		},
		{
			operation: "change user number",
			command:   "/change_user_number",
			expect:    event.ChangeUserNumber,
		},
		{
			operation: "start voting",
			command:   "/start_voting",
			expect:    event.StartVoting,
		},
		{
			operation: "end session",
			command:   "/end_session",
			expect:    event.EndSession,
		},
		{
			operation: "unknown operation",
			command:   "/unknown_command",
			expect:    event.Unknown,
		},
	} {
		t.Run(testCase.operation, func(t *testing.T) {
			t.Parallel()

			u := &tgbotapi.Update{
				Message: &tgbotapi.Message{
					Text: testCase.command,
				},
			}

			e := event.From(u)

			assert.Equal(t, testCase.expect, e.Command())
		})
	}
}

func TestEvent_UserID(t *testing.T) {
	t.Parallel()

	t.Run("no user id when there is no message", func(t *testing.T) {
		t.Parallel()

		u := &tgbotapi.Update{
			Message: nil,
		}

		e := event.From(u)

		assert.Empty(t, e.UserID())
	})

	t.Run("no user id when there is no user", func(t *testing.T) {
		t.Parallel()

		u := &tgbotapi.Update{
			Message: &tgbotapi.Message{
				From: nil,
			},
		}

		e := event.From(u)

		assert.Empty(t, e.UserID())
	})

	t.Run("user id", func(t *testing.T) {
		t.Parallel()

		u := &tgbotapi.Update{
			Message: &tgbotapi.Message{
				From: &tgbotapi.User{
					ID: 111111111,
				},
			},
		}

		e := event.From(u)

		assert.Equal(t, "111111111", e.UserID())
	})
}

func TestEvent_Payload(t *testing.T) {
	t.Parallel()

	t.Run("empty if no message", func(t *testing.T) {
		t.Parallel()

		u := &tgbotapi.Update{
			Message: nil,
		}

		e := event.From(u)

		assert.Empty(t, e.Payload())
	})

	t.Run("is message text", func(t *testing.T) {
		t.Parallel()

		u := &tgbotapi.Update{
			Message: &tgbotapi.Message{
				Text: "some text",
			},
		}

		e := event.From(u)

		assert.Equal(t, "some text", e.Payload())
	})
}
