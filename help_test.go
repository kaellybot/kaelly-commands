package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedHelpCustomID = "/help"

	expectedHelpCommandName  = "item"
	expectedHelpPageCustomID = "/help/item/page"
)

func TestCraftHelpCustomID(t *testing.T) {
	actual := commands.CraftHelpCustomID()
	assert.Equal(t, expectedHelpCustomID, actual,
		"CraftHelpCustomID should return the expected Item custom ID")
}

func TestCraftHelpPageCustomID(t *testing.T) {
	actual := commands.CraftHelpPageCustomID(expectedHelpCommandName)
	assert.Equal(t, expectedHelpPageCustomID, actual,
		"CraftHelpPageCustomID should return the expected Page custom ID")
}

func TestExtractHelpPageCustomID(t *testing.T) {
	tests := []struct {
		name            string
		customID        string
		expectedCommand string
		succeeded       bool
	}{
		{
			name:     "HelpPageCustomID could not be extracted",
			customID: expectedHelpCustomID,
		},
		{
			name:            "HelpPageCustomID nominal case",
			customID:        expectedHelpPageCustomID,
			expectedCommand: expectedHelpCommandName,
			succeeded:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandName, ok := commands.ExtractHelpPageCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCommand, commandName)
			} else {
				assert.False(t, ok)
			}
		})
	}
}
