package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedHelpCustomID = "/help"

	expectedHelpCommandName  = "item"
	expectedHelpPageCustomID = "/help/item/page"
)

func TestCraftHelpCustomID(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "CraftHelpCustomID returns expected help custom ID",
			expected: expectedHelpCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CraftHelpCustomID()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCraftHelpPageCustomID(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		expected string
	}{
		{
			name:     "CraftHelpPageCustomID returns expected help page custom ID",
			command:  expectedHelpCommandName,
			expected: expectedHelpPageCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CraftHelpPageCustomID(tt.command)
			assert.Equal(t, tt.expected, actual)
		})
	}
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
			commandName, ok := ExtractHelpPageCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCommand, commandName)
			} else {
				assert.False(t, ok)
			}
		})
	}
}
