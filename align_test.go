package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedAlignBookPageCustomID        = "/books/align?city=bonta&order=esprit&server=draconiros&page=2"
	expectedAlignBookPageNoCityCustomID  = "/books/align?city=_&order=esprit&server=draconiros&page=2"
	expectedAlignBookPageNoOrderCustomID = "/books/align?city=bonta&order=_&server=draconiros&page=2"
	expectedAlignBookPageNoAllCustomID   = "/books/align?city=_&order=_&server=draconiros&page=2"

	expectedAlignBookCityNoOrderCustomID = "/books/align?order=_&server=draconiros"
	expectedAlignBookCityCustomID        = "/books/align?order=esprit&server=draconiros"

	expectedAlignBookOrderNoCityCustomID = "/books/align?city=_&server=draconiros"
	expectedAlignBookOrderCustomID       = "/books/align?city=bonta&server=draconiros"

	expectedAlignBookPage   = 2
	expectedAlignBookCity   = "bonta"
	expectedAlignBookOrder  = "esprit"
	expectedAlignBookServer = "draconiros"
)

func TestCraftAlignBookPageCustomID(t *testing.T) {
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
			expected: expectedAlignBookPageNoCityCustomID,
		},
		{
			name:     "CraftAlignBookCustomID without order returns expected custom ID",
			city:     expectedAlignBookCity,
			order:    "",
			server:   expectedAlignBookServer,
			page:     expectedAlignBookPage,
			expected: expectedAlignBookPageNoOrderCustomID,
		},
		{
			name:     "CraftAlignBookCustomID without order and city returns expected custom ID",
			city:     "",
			order:    "",
			server:   expectedAlignBookServer,
			page:     expectedAlignBookPage,
			expected: expectedAlignBookPageNoAllCustomID,
		},
		{
			name:     "CraftAlignBookCustomID returns expected custom ID",
			city:     expectedAlignBookCity,
			order:    expectedAlignBookOrder,
			server:   expectedAlignBookServer,
			page:     expectedAlignBookPage,
			expected: expectedAlignBookPageCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftAlignBookPageCustomID(tt.city, tt.order, tt.server, tt.page)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractAlignBookPageCustomID(t *testing.T) {
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
			name:     "AlignBookCityCustomID could not be converted",
			customID: expectedAlignBookCityCustomID,
		},
		{
			name:     "AlignBookOrderCustomID could not be converted",
			customID: expectedAlignBookOrderCustomID,
		},
		{
			name:           "AlignBookCustomID (without city) nominal case",
			customID:       expectedAlignBookPageNoCityCustomID,
			expectedCity:   "",
			expectedOrder:  expectedAlignBookOrder,
			expectedServer: expectedAlignBookServer,
			expectedPage:   expectedAlignBookPage,
			succeeded:      true,
		},
		{
			name:           "AlignBookCustomID (without order) nominal case",
			customID:       expectedAlignBookPageNoOrderCustomID,
			expectedCity:   expectedAlignBookCity,
			expectedOrder:  "",
			expectedServer: expectedAlignBookServer,
			expectedPage:   expectedAlignBookPage,
			succeeded:      true,
		},
		{
			name:           "AlignBookCustomID (without city and order) nominal case",
			customID:       expectedAlignBookPageNoAllCustomID,
			expectedCity:   "",
			expectedOrder:  "",
			expectedServer: expectedAlignBookServer,
			expectedPage:   expectedAlignBookPage,
			succeeded:      true,
		},
		{
			name:           "AlignBookCustomID nominal case",
			customID:       expectedAlignBookPageCustomID,
			expectedCity:   expectedAlignBookCity,
			expectedOrder:  expectedAlignBookOrder,
			expectedServer: expectedAlignBookServer,
			expectedPage:   expectedAlignBookPage,
			succeeded:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			city, order, server, page, ok := commands.ExtractAlignBookPageCustomID(tt.customID)
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

func TestCraftAlignBookCityCustomID(t *testing.T) {
	tests := []struct {
		name     string
		order    string
		server   string
		expected string
	}{
		{
			name:     "CraftAlignBookCityCustomID without order returns expected custom ID",
			order:    "",
			server:   expectedAlignBookServer,
			expected: expectedAlignBookCityNoOrderCustomID,
		},
		{
			name:     "CraftAlignBookCustomID returns expected custom ID",
			order:    expectedAlignBookOrder,
			server:   expectedAlignBookServer,
			expected: expectedAlignBookCityCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftAlignBookCityCustomID(tt.order, tt.server)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractAlignBookCityCustomID(t *testing.T) {
	tests := []struct {
		name           string
		customID       string
		expectedOrder  string
		expectedServer string
		succeeded      bool
	}{
		{
			name:     "AlignBookCustomID could not be extracted",
			customID: "/other/id",
		},
		{
			name:     "AlignBookPageCustomID could not be converted",
			customID: expectedAlignBookPageCustomID,
		},
		{
			name:     "AlignBookOrderCustomID could not be converted",
			customID: expectedAlignBookOrderCustomID,
		},
		{
			name:           "AlignBookCityCustomID (without order) nominal case",
			customID:       expectedAlignBookCityNoOrderCustomID,
			expectedOrder:  "",
			expectedServer: expectedAlignBookServer,
			succeeded:      true,
		},
		{
			name:           "AlignBookCustomID nominal case",
			customID:       expectedAlignBookCityCustomID,
			expectedOrder:  expectedAlignBookOrder,
			expectedServer: expectedAlignBookServer,
			succeeded:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, server, ok := commands.ExtractAlignBookCityCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedOrder, order)
				assert.Equal(t, tt.expectedServer, server)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestCraftAlignBookOrderCustomID(t *testing.T) {
	tests := []struct {
		name     string
		city     string
		server   string
		expected string
	}{
		{
			name:     "CraftAlignBookOrderCustomID without city returns expected custom ID",
			city:     "",
			server:   expectedAlignBookServer,
			expected: expectedAlignBookOrderNoCityCustomID,
		},
		{
			name:     "CraftAlignBookOrderCustomID returns expected custom ID",
			city:     expectedAlignBookCity,
			server:   expectedAlignBookServer,
			expected: expectedAlignBookOrderCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftAlignBookOrderCustomID(tt.city, tt.server)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractAlignBookOrderCustomID(t *testing.T) {
	tests := []struct {
		name           string
		customID       string
		expectedCity   string
		expectedServer string
		succeeded      bool
	}{
		{
			name:     "AlignBookCustomID could not be extracted",
			customID: "/other/id",
		},
		{
			name:     "AlignBookPageCustomID could not be converted",
			customID: expectedAlignBookPageCustomID,
		},
		{
			name:     "AlignBookCityCustomID could not be converted",
			customID: expectedAlignBookCityCustomID,
		},
		{
			name:           "AlignBookOrderCustomID (without city) nominal case",
			customID:       expectedAlignBookOrderNoCityCustomID,
			expectedCity:   "",
			expectedServer: expectedAlignBookServer,
			succeeded:      true,
		},
		{
			name:           "AlignBookOrderCustomID nominal case",
			customID:       expectedAlignBookOrderCustomID,
			expectedCity:   expectedAlignBookCity,
			expectedServer: expectedAlignBookServer,
			succeeded:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			city, server, ok := commands.ExtractAlignBookOrderCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCity, city)
				assert.Equal(t, tt.expectedServer, server)
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
			input:    expectedAlignBookPageCustomID,
			expected: true,
		},
		{
			name:     "Valid align book (without city) custom ID",
			input:    expectedAlignBookPageNoCityCustomID,
			expected: true,
		},
		{
			name:     "Valid align book (without order) custom ID",
			input:    expectedAlignBookPageNoOrderCustomID,
			expected: true,
		},
		{
			name:     "Valid align book (without city and order) custom ID",
			input:    expectedAlignBookPageNoAllCustomID,
			expected: true,
		},
		{
			name:     "Valid align book city (without order) custom ID",
			input:    expectedAlignBookCityNoOrderCustomID,
			expected: true,
		},
		{
			name:     "Valid align book city custom ID",
			input:    expectedAlignBookCityCustomID,
			expected: true,
		},
		{
			name:     "Valid align book order (without city) custom ID",
			input:    expectedAlignBookOrderNoCityCustomID,
			expected: true,
		},
		{
			name:     "Valid align book order custom ID",
			input:    expectedAlignBookOrderCustomID,
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
