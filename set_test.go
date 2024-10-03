package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedSetID            = "123"
	expectedSetCustomID      = "/set/123"
	expectedSetBonusCustomID = "/set/123/bonuses"
)

func TestCraftSetCustomID(t *testing.T) {
	tests := []struct {
		name     string
		setID    string
		expected string
	}{
		{
			name:     "CraftSetCustomID returns expected set custom ID",
			setID:    expectedSetID,
			expected: expectedSetCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftSetCustomID(tt.setID)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCraftSetBonusCustomID(t *testing.T) {
	tests := []struct {
		name     string
		setID    string
		expected string
	}{
		{
			name:     "CraftSetBonusCustomID returns expected bonus custom ID",
			setID:    expectedSetID,
			expected: expectedSetBonusCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.CraftSetBonusCustomID(tt.setID)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractSetCustomID(t *testing.T) {
	tests := []struct {
		name     string
		customID string
		expected string
		ok       bool
	}{
		{
			name:     "ExtractSetCustomID successfully extracts set ID",
			customID: expectedSetCustomID,
			expected: expectedSetID,
			ok:       true,
		},
		{
			name:     "ExtractSetCustomID fails on bonus custom ID",
			customID: expectedSetBonusCustomID,
			ok:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualSetID, ok := commands.ExtractSetCustomID(tt.customID)
			assert.Equal(t, tt.ok, ok)
			if tt.ok {
				assert.Equal(t, tt.expected, actualSetID)
			}
		})
	}
}

func TestExtractSetBonusCustomID(t *testing.T) {
	tests := []struct {
		name     string
		customID string
		expected string
		ok       bool
	}{
		{
			name:     "ExtractSetBonusCustomID successfully extracts set ID",
			customID: expectedSetBonusCustomID,
			expected: expectedSetID,
			ok:       true,
		},
		{
			name:     "ExtractSetBonusCustomID fails on set custom ID",
			customID: expectedSetCustomID,
			ok:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualSetID, ok := commands.ExtractSetBonusCustomID(tt.customID)
			assert.Equal(t, tt.ok, ok)
			if tt.ok {
				assert.Equal(t, tt.expected, actualSetID)
			}
		})
	}
}

func TestIsBelongsToSet(t *testing.T) {
	tests := []struct {
		name     string
		customID string
		expected bool
	}{
		{
			name:     "IsBelongsToSet returns true for valid set custom ID",
			customID: expectedSetCustomID,
			expected: true,
		},
		{
			name:     "IsBelongsToSet returns true for valid bonus custom ID",
			customID: expectedSetBonusCustomID,
			expected: true,
		},
		{
			name:     "IsBelongsToSet returns false for invalid custom ID",
			customID: "/other/123",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := commands.IsBelongsToSet(tt.customID)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
