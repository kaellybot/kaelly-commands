package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedAlignBookCustomID        = "/align/book?city=bonta&order=esprit&server=draconiros&page=2"
	expectedAlignBookNoCityCustomID  = "/align/book?city=_&order=esprit&server=draconiros&page=2"
	expectedAlignBookNoOrderCustomID = "/align/book?city=bonta&order=_&server=draconiros&page=2"
	expectedAlignBookNoAllCustomID   = "/align/book?city=_&order=_&server=draconiros&page=2"

	expectedAlignBookPage   = 2
	expectedAlignBookCity   = "bonta"
	expectedAlignBookOrder  = "esprit"
	expectedAlignBookServer = "draconiros"
)

func TestCraftAlignBookCustomID(t *testing.T) {
	tests := []struct {
		name     string
		city     string
		order    string
		server   string
		expected string
		page     int
	}{
		{
			name:     "CraftAlignBookCustomID without city returns expected custom ID",
			city:     "",
			order:    expectedAlignBookOrder,
			server:   expectedAlignBookServer,
			page:     expectedAlignBookPage,
			expected: expectedAlignBookNoCityCustomID,
		},
		{
			name:     "CraftAlignBookCustomID without order returns expected custom ID",
			city:     expectedAlignBookCity,
			order:    "",
			server:   expectedAlignBookServer,
			page:     expectedAlignBookPage,
			expected: expectedAlignBookNoOrderCustomID,
		},
		{
			name:     "CraftAlignBookCustomID without order and city returns expected custom ID",
			city:     "",
			order:    "",
			server:   expectedAlignBookServer,
			page:     expectedAlignBookPage,
			expected: expectedAlignBookNoAllCustomID,
		},
		{
			name:     "CraftAlignBookCustomID returns expected custom ID",
			city:     expectedAlignBookCity,
			order:    expectedAlignBookOrder,
			server:   expectedAlignBookServer,
			page:     expectedAlignBookPage,
			expected: expectedAlignBookCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftAlignBookCustomID(tt.city, tt.order, tt.server, tt.page)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractAlignBookCustomID(t *testing.T) {
	tests := []struct {
		name           string
		customID       string
		expectedCity   string
		expectedOrder  string
		expectedServer string
		expectedPage   int
		succeeded      bool
	}{
		{
			name:     "AlignBookCustomID could not be extracted",
			customID: "/other/id",
		},
		{
			name:     "AlignBookCustomID could not be converted",
			customID: "/align/book?city=_&order=_&server=draconiros&page=9999999999999999999",
		},
		{
			name:           "AlignBookCustomID (without city) nominal case",
			customID:       expectedAlignBookNoCityCustomID,
			expectedCity:   "",
			expectedOrder:  expectedAlignBookOrder,
			expectedServer: expectedAlignBookServer,
			expectedPage:   expectedAlignBookPage,
			succeeded:      true,
		},
		{
			name:           "AlignBookCustomID (without order) nominal case",
			customID:       expectedAlignBookNoOrderCustomID,
			expectedCity:   expectedAlignBookCity,
			expectedOrder:  "",
			expectedServer: expectedAlignBookServer,
			expectedPage:   expectedAlignBookPage,
			succeeded:      true,
		},
		{
			name:           "AlignBookCustomID (without city and order) nominal case",
			customID:       expectedAlignBookNoAllCustomID,
			expectedCity:   "",
			expectedOrder:  "",
			expectedServer: expectedAlignBookServer,
			expectedPage:   expectedAlignBookPage,
			succeeded:      true,
		},
		{
			name:           "AlignBookCustomID nominal case",
			customID:       expectedAlignBookCustomID,
			expectedCity:   expectedAlignBookCity,
			expectedOrder:  expectedAlignBookOrder,
			expectedServer: expectedAlignBookServer,
			expectedPage:   expectedAlignBookPage,
			succeeded:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			city, order, server, page, ok := commands.ExtractAlignBookCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCity, city)
				assert.Equal(t, tt.expectedOrder, order)
				assert.Equal(t, tt.expectedServer, server)
				assert.Equal(t, tt.expectedPage, page)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestIsBelongsToAlign(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid align book custom ID",
			input:    expectedAlignBookCustomID,
			expected: true,
		},
		{
			name:     "Valid align book (without city) custom ID",
			input:    expectedAlignBookNoCityCustomID,
			expected: true,
		},
		{
			name:     "Valid align book (without order) custom ID",
			input:    expectedAlignBookNoOrderCustomID,
			expected: true,
		},
		{
			name:     "Valid align book (without city and order) custom ID",
			input:    expectedAlignBookNoAllCustomID,
			expected: true,
		},
		{
			name:  "Invalid custom ID",
			input: "/other/123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, commands.IsBelongsToAlign(tt.input))
		})
	}
}
