package commands_test

import (
	"testing"
	"time"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedAlmanaxDayCustomID               = "/almanax/day/761702400"
	expectedAlmanaxResourceCharacterCustomID = "/almanax/resource?duration=30"
	expectedAlmanaxResourceDurationCustomID  = "/almanax/resource?characters=8"

	expectedCharacters = 8
	expectedDuration   = 30
)

func getExpectedDate() time.Time {
	return time.Date(1994, time.February, 20, 0, 0, 0, 0, time.UTC)
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
			customID: expectedAlmanaxResourceCharacterCustomID,
		},
		{
			name:     "AlmanaxDayCustomID could not be extracted",
			customID: expectedAlmanaxResourceDurationCustomID,
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

func TestCraftAlmanaxResourceCharacterCustomID(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		duration int64
	}{
		{
			name:     "CraftAlmanaxResourceCharacterCustomID returns expected custom ID",
			duration: expectedDuration,
			expected: expectedAlmanaxResourceCharacterCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftAlmanaxResourceCharacterCustomID(tt.duration)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractAlmanaxResourceCharacterCustomID(t *testing.T) {
	tests := []struct {
		name             string
		customID         string
		expectedDuration int64
		succeeded        bool
	}{
		{
			name:     "AlmanaxResourceCharacterCustomID could not be extracted",
			customID: expectedAlmanaxDayCustomID,
		},
		{
			name:     "AlmanaxResourceCharacterCustomID could not be extracted",
			customID: expectedAlmanaxResourceDurationCustomID,
		},
		{
			name:     "AlmanaxResourceCharacterCustomID could not be converted",
			customID: "/almanax/resource?duration=9999999999999999999",
		},
		{
			name:             "AlmanaxResourceCharacterCustomID nominal case",
			customID:         expectedAlmanaxResourceCharacterCustomID,
			expectedDuration: expectedDuration,
			succeeded:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration, ok := commands.ExtractAlmanaxResourceCharacterCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedDuration, duration)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestCraftAlmanaxResourceDurationCustomID(t *testing.T) {
	tests := []struct {
		name            string
		expected        string
		characterNumber int64
	}{
		{
			name:            "CraftAlmanaxResourceDurationCustomID returns expected custom ID",
			characterNumber: expectedCharacters,
			expected:        expectedAlmanaxResourceDurationCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftAlmanaxResourceDurationCustomID(tt.characterNumber)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractAlmanaxResourceDurationCustomID(t *testing.T) {
	tests := []struct {
		name               string
		customID           string
		expectedCharacters int64
		succeeded          bool
	}{
		{
			name:     "AlmanaxResourceDurationCustomID could not be extracted",
			customID: expectedAlmanaxDayCustomID,
		},
		{
			name:     "AlmanaxResourceDurationCustomID could not be extracted",
			customID: expectedAlmanaxResourceCharacterCustomID,
		},
		{
			name:     "AlmanaxResourceDurationCustomID could not be converted",
			customID: "/almanax/resource?characters=9999999999999999999",
		},
		{
			name:               "AlmanaxResourceDurationCustomID nominal case",
			customID:           expectedAlmanaxResourceDurationCustomID,
			expectedCharacters: expectedCharacters,
			succeeded:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration, ok := commands.ExtractAlmanaxResourceDurationCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedCharacters, duration)
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
			input:    expectedAlmanaxResourceCharacterCustomID,
			expected: true,
		},
		{
			name:     "Valid almanax resource custom ID",
			input:    expectedAlmanaxResourceDurationCustomID,
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
