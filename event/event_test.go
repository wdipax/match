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
			operation: "add team member",
			command:   "/add_team_member",
			expect:    event.AddTeamMember,
		},
		{
			operation: "team member name",
			command:   "/team_member_name",
			expect:    event.TeamMemberName,
		},
		{
			operation: "team member number",
			command:   "/team_member_number",
			expect:    event.TeamMemberNumber,
		},
		{
			operation: "start voting",
			command:   "/start_voting",
			expect:    event.StartVoting,
		},
		{
			operation: "vote",
			command:   "/vote",
			expect:    event.Vote,
		},
		{
			operation: "end session",
			command:   "/end_session",
			expect:    event.EndSession,
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
