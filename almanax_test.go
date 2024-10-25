package commands_test

import (
	"testing"
	"time"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedAlmanaxDayCustomID      = "/almanax/day/761702400"
	expectedAlmanaxResourceCustomID = "/almanax/resource?startDate=761702400&endDate=761788800"
)

func getExpectedDate() time.Time {
	return time.Date(1994, time.February, 20, 0, 0, 0, 0, time.UTC)
}

func getExpectedEndDate() time.Time {
	return time.Date(1994, time.February, 21, 0, 0, 0, 0, time.UTC)
}

func TestCraftAlmanaxDayCustomID(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "CraftAlmanaxDayCustomID returns expected time custom ID",
			date:     getExpectedDate(),
			expected: expectedAlmanaxDayCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftAlmanaxDayCustomID(tt.date)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractAlmanaxDayCustomID(t *testing.T) {
	expectedDate := getExpectedDate()
	tests := []struct {
		name         string
		customID     string
		expectedDate *time.Time
		succeeded    bool
	}{
		{
			name:     "AlmanaxDayCustomID could not be extracted",
			customID: expectedAlmanaxResourceCustomID,
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
			date, ok := commands.ExtractAlmanaxDayCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedDate, date)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestCraftAlmanaxResourceCustomID(t *testing.T) {
	tests := []struct {
		name      string
		expected  string
		startDate time.Time
		endDate   time.Time
	}{
		{
			name:      "CraftAlmanaxResourceCustomID returns expected time custom ID",
			startDate: getExpectedDate(),
			endDate:   getExpectedEndDate(),
			expected:  expectedAlmanaxResourceCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftAlmanaxResourceCustomID(tt.startDate, tt.endDate)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractAlmanaxResourceCustomID(t *testing.T) {
	expectedDate := getExpectedDate()
	expectedEndDate := getExpectedEndDate()
	tests := []struct {
		name              string
		customID          string
		expectedStartDate *time.Time
		expectedEndDate   *time.Time
		succeeded         bool
	}{
		{
			name:     "AlmanaxResourceCustomID could not be extracted",
			customID: expectedAlmanaxDayCustomID,
		},
		{
			name:     "AlmanaxResourceCustomID could not be converted",
			customID: "/almanax/resource?startDate=9999999999999999999&endDate=123",
		},
		{
			name:     "AlmanaxResourceCustomID could not be converted 2",
			customID: "/almanax/resource?startDate=123&endDate=9999999999999999999",
		},
		{
			name:              "AlmanaxResourceCustomID nominal case",
			customID:          expectedAlmanaxResourceCustomID,
			expectedStartDate: &expectedDate,
			expectedEndDate:   &expectedEndDate,
			succeeded:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startDate, endDate, ok := commands.ExtractAlmanaxResourceCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedStartDate, startDate)
				assert.Equal(t, tt.expectedEndDate, endDate)
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
			name:     "Valid almanax resource custom ID",
			input:    expectedAlmanaxResourceCustomID,
			expected: true,
		},
		{
			name:  "Invalid custom ID",
			input: "/other/123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, commands.IsBelongsToAlmanax(tt.input))
		})
	}
}
