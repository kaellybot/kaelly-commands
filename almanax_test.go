package commands

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	expectedAlmanaxDayCustomID = "/almanax/day/761702400"
)

var (
	expectedDate = time.Date(1994, time.February, 20, 0, 0, 0, 0, time.UTC)
)

func TestCraftAlmanaxDayCustomID(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "CraftAlmanaxDayCustomID returns expected time custom ID",
			date:     expectedDate,
			expected: expectedAlmanaxDayCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CraftAlmanaxDayCustomID(tt.date)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractAlmanaxDayCustomID(t *testing.T) {
	tests := []struct {
		name         string
		customID     string
		expectedDate *time.Time
		succeeded    bool
	}{
		{
			name:     "AlmanaxDayCustomID could not be extracted",
			customID: "/other/id/49849844",
		},
		{
			name:     "AlmanaxDayCustomID could not be converted",
			customID: "/almanax/day/9999999999999999999",
		},
		{
			name:         "AlmanaxDayCustomID nominal case",
			customID:     expectedAlmanaxDayCustomID,
			expectedDate: &expectedDate,
			succeeded:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date, ok := ExtractAlmanaxDayCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedDate, date)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestIsBelongsToAlmanax(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid almanax day custom ID",
			input:    expectedAlmanaxDayCustomID,
			expected: true,
		},
		{
			name:  "Invalid custom ID",
			input: "/other/123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsBelongsToAlmanax(tt.input))
		})
	}
}
