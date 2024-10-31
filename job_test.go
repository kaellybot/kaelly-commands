package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedJobBookCustomID = "/job/book?job=bucheron&server=draconiros&page=2"

	expectedJobBookPage   = 2
	expectedJobBookJob    = "bucheron"
	expectedJobBookServer = "draconiros"
)

func TestCraftJobBookCustomID(t *testing.T) {
	tests := []struct {
		name     string
		job      string
		server   string
		expected string
		page     int
	}{
		{
			name:     "CraftJobBookCustomID returns expected custom ID",
			job:      expectedJobBookJob,
			server:   expectedJobBookServer,
			page:     expectedJobBookPage,
			expected: expectedJobBookCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftJobBookCustomID(tt.job, tt.server, tt.page)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractJobBookCustomID(t *testing.T) {
	tests := []struct {
		name           string
		customID       string
		expectedJob    string
		expectedServer string
		expectedPage   int
		succeeded      bool
	}{
		{
			name:     "JobBookCustomID could not be extracted",
			customID: "/other/id",
		},
		{
			name:     "JobBookCustomID could not be converted",
			customID: "/job/book?job=bucheron&server=draconiros&page=9999999999999999999",
		},
		{
			name:           "JobBookCustomID nominal case",
			customID:       expectedJobBookCustomID,
			expectedJob:    expectedJobBookJob,
			expectedServer: expectedJobBookServer,
			expectedPage:   expectedJobBookPage,
			succeeded:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job, server, page, ok := commands.ExtractJobBookCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedJob, job)
				assert.Equal(t, tt.expectedServer, server)
				assert.Equal(t, tt.expectedPage, page)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestIsBelongsToJob(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid job book custom ID",
			input:    expectedJobBookCustomID,
			expected: true,
		},
		{
			name:  "Invalid custom ID",
			input: "/other/123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, commands.IsBelongsToJob(tt.input))
		})
	}
}
