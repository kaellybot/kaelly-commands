package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedJobBookPageCustomID   = "/books/job/bucheron?server=draconiros&page=2"
	expectedJobBookSelectCustomID = "/books/job?server=draconiros"

	expectedJobBookPage   = 2
	expectedJobBookJob    = "bucheron"
	expectedJobBookServer = "draconiros"
)

func TestCraftJobBookPageCustomID(t *testing.T) {
	tests := []struct {
		name     string
		job      string
		server   string
		expected string
		page     int
	}{
		{
			name:     "CraftJobBookPageCustomID returns expected custom ID",
			job:      expectedJobBookJob,
			server:   expectedJobBookServer,
			page:     expectedJobBookPage,
			expected: expectedJobBookPageCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftJobBookPageCustomID(tt.job, tt.server, tt.page)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractJobBookPageCustomID(t *testing.T) {
	tests := []struct {
		name           string
		customID       string
		expectedJob    string
		expectedServer string
		expectedPage   int
		succeeded      bool
	}{
		{
			name:     "JobBookPageCustomID could not be extracted",
			customID: "/other/id",
		},
		{
			name:     "JobBookPageCustomID could not be converted",
			customID: "/books/job?server=draconiros",
		},
		{
			name:     "JobBookPageCustomID could not be converted",
			customID: "/books/job/bucheron?server=draconiros&page=9999999999999999999",
		},
		{
			name:           "JobBookPageCustomID nominal case",
			customID:       expectedJobBookPageCustomID,
			expectedJob:    expectedJobBookJob,
			expectedServer: expectedJobBookServer,
			expectedPage:   expectedJobBookPage,
			succeeded:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			job, server, page, ok := commands.ExtractJobBookPageCustomID(tt.customID)
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

func TestCraftJobBookSelectCustomID(t *testing.T) {
	tests := []struct {
		name     string
		server   string
		expected string
	}{
		{
			name:     "CraftJobBookSelectCustomID returns expected custom ID",
			server:   expectedJobBookServer,
			expected: expectedJobBookSelectCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftJobBookSelectCustomID(tt.server)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractJobBookSelectCustomID(t *testing.T) {
	tests := []struct {
		name           string
		customID       string
		expectedServer string
		succeeded      bool
	}{
		{
			name:     "JobBookPageCustomID could not be extracted",
			customID: "/other/id",
		},
		{
			name:     "JobBookPageCustomID could not be converted",
			customID: "/books/job?server=draconiros&page=1",
		},
		{
			name:           "JobBookSelectCustomID nominal case",
			customID:       expectedJobBookSelectCustomID,
			expectedServer: expectedJobBookServer,
			succeeded:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, ok := commands.ExtractJobBookSelectCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedServer, server)
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
			name:     "Valid job book page custom ID",
			input:    expectedJobBookPageCustomID,
			expected: true,
		},
		{
			name:     "Valid job book select custom ID",
			input:    expectedJobBookSelectCustomID,
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
