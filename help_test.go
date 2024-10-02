package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedHelpCustomID = "/help"

	expectedHelpCommandName  = "item"
	expectedHelpPage         = 2
	expectedHelpPageCustomID = "/help/item/details?page=2"
)

func TestCraftHelpCustomID(t *testing.T) {
	actual := commands.CraftHelpCustomID()
	assert.Equal(t, expectedHelpCustomID, actual,
		"CraftHelpCustomID should return the expected Item custom ID")
}

func TestCraftHelpPageCustomID(t *testing.T) {
	actual := commands.CraftHelpPageCustomID(expectedHelpCommandName, expectedHelpPage)
	assert.Equal(t, expectedHelpPageCustomID, actual,
		"CraftHelpPageCustomID should return the expected Page custom ID")
}

func TestExtractHelpPageCustomID(t *testing.T) {
	tests := []struct {
		name            string
		customID        string
		expectedCommand string
		expectedPage    int
		succeeded       bool
	}{
		{
			name:     "HelpPageCustomID could not be extracted",
			customID: expectedHelpCustomID,
		},
		{
			name:     "HelpPageCustomID page could not be parsed",
			customID: "/help/item/details?page=9999999999999999999",
		},
		{
			name:            "HelpPageCustomID nominal case",
			customID:        expectedHelpPageCustomID,
			expectedCommand: expectedHelpCommandName,
			expectedPage:    expectedHelpPage,
			succeeded:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandName, page, ok := commands.ExtractHelpPageCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCommand, commandName)
				assert.Equal(t, tt.expectedPage, page)
			} else {
				assert.False(t, ok)
			}
		})
	}
}
