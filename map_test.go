package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedMapNormalCustomID   = "/map/42?type=normal"
	expectedMapTacticalCustomID = "/map/42?type=tactical"

	expectedMapNumber = 42
)

func TestCraftMapNormalCustomID(t *testing.T) {
	tests := []struct {
		name      string
		mapNumber int64
		expected  string
	}{
		{
			name:      "CraftMapNormalCustomID returns expected map normal custom ID",
			mapNumber: expectedMapNumber,
			expected:  expectedMapNormalCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CraftMapNormalCustomID(tt.mapNumber)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCraftMapTacticalCustomID(t *testing.T) {
	tests := []struct {
		name      string
		mapNumber int64
		expected  string
	}{
		{
			name:      "CraftMapTacticalCustomID returns expected map tactical custom ID",
			mapNumber: expectedMapNumber,
			expected:  expectedMapTacticalCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CraftMapTacticalCustomID(tt.mapNumber)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractMapNormalCustomID(t *testing.T) {
	tests := []struct {
		name              string
		customID          string
		expectedMapNumber int64
		succeeded         bool
	}{
		{
			name:     "MapNormalCustomID could not be extracted",
			customID: expectedMapTacticalCustomID,
		},
		{
			name:     "MapNormalCustomID could not be converted",
			customID: "/map/9999999999999999999?type=normal",
		},
		{
			name:              "MapNormalCustomID nominal case",
			customID:          expectedMapNormalCustomID,
			expectedMapNumber: expectedMapNumber,
			succeeded:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapNumber, ok := ExtractMapNormalCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedMapNumber, mapNumber)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestExtractMapTacticalCustomID(t *testing.T) {
	tests := []struct {
		name              string
		customID          string
		expectedMapNumber int64
		succeeded         bool
	}{
		{
			name:     "MapTacticalCustomID could not be extracted",
			customID: expectedMapNormalCustomID,
		},
		{
			name:     "MapTacticalCustomID could not be converted",
			customID: "/map/9999999999999999999?type=tactical",
		},
		{
			name:              "MapTacticalCustomID nominal case",
			customID:          expectedMapTacticalCustomID,
			expectedMapNumber: expectedMapNumber,
			succeeded:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapNumber, ok := ExtractMapTacticalCustomID(tt.customID)
			if tt.succeeded {
				assert.True(t, ok)
				assert.Equal(t, tt.expectedMapNumber, mapNumber)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestIsBelongsToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid map normal custom ID",
			input:    expectedMapNormalCustomID,
			expected: true,
		},
		{
			name:     "Valid map tactical custom ID",
			input:    expectedMapTacticalCustomID,
			expected: true,
		},
		{
			name:  "Invalid custom ID",
			input: "/other/123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsBelongsToMap(tt.input))
		})
	}
}
