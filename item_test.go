package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedItemID              = "123"
	expectedItemType            = "QUEST_ITEM"
	expectedItemCustomID        = "/item?type=QUEST_ITEM"
	expectedItemEffectsCustomID = "/item/123/effects?type=QUEST_ITEM"
	expectedItemRecipeCustomID  = "/item/123/recipe?type=QUEST_ITEM"
)

func TestCraftItemCustomID(t *testing.T) {
	tests := []struct {
		name     string
		itemType string
		expected string
	}{
		{
			name:     "CraftItemCustomID returns expected custom ID",
			itemType: expectedItemType,
			expected: expectedItemCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CraftItemCustomID(tt.itemType)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCraftItemEffectsCustomID(t *testing.T) {
	tests := []struct {
		name     string
		itemID   string
		itemType string
		expected string
	}{
		{
			name:     "CraftItemEffectsCustomID returns expected effects custom ID",
			itemID:   expectedItemID,
			itemType: expectedItemType,
			expected: expectedItemEffectsCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CraftItemEffectsCustomID(tt.itemID, tt.itemType)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCraftItemRecipeCustomID(t *testing.T) {
	tests := []struct {
		name     string
		itemID   string
		itemType string
		expected string
	}{
		{
			name:     "CraftItemRecipeCustomID returns expected recipe custom ID",
			itemID:   expectedItemID,
			itemType: expectedItemType,
			expected: expectedItemRecipeCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CraftItemRecipeCustomID(tt.itemID, tt.itemType)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractItemCustomID(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedType   string
		expectedResult bool
	}{
		{
			name:           "Nominal case",
			input:          expectedItemCustomID,
			expectedType:   expectedItemType,
			expectedResult: true,
		},
		{
			name:  "Bad case (Effects custom ID)",
			input: expectedItemEffectsCustomID,
		},
		{
			name:  "Bad case (Recipe custom ID)",
			input: expectedItemRecipeCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualItemType, ok := ExtractItemCustomID(tt.input)
			assert.Equal(t, tt.expectedResult, ok)
			if ok {
				assert.Equal(t, tt.expectedType, actualItemType)
			}
		})
	}
}

func TestExtractItemEffectsCustomID(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedID     string
		expectedType   string
		expectedResult bool
	}{
		{
			name:           "Nominal case",
			input:          expectedItemEffectsCustomID,
			expectedID:     expectedItemID,
			expectedType:   expectedItemType,
			expectedResult: true,
		},
		{
			name:  "Bad case (Custom ID)",
			input: expectedItemCustomID,
		},
		{
			name:  "Bad case (Recipe custom ID)",
			input: expectedItemRecipeCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualItemID, actualItemType, ok := ExtractItemEffectsCustomID(tt.input)
			assert.Equal(t, tt.expectedResult, ok)
			if ok {
				assert.Equal(t, tt.expectedID, actualItemID)
				assert.Equal(t, tt.expectedType, actualItemType)
			}
		})
	}
}

func TestExtractItemRecipeCustomID(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedID     string
		expectedType   string
		expectedResult bool
	}{
		{
			name:           "Nominal case",
			input:          expectedItemRecipeCustomID,
			expectedID:     expectedItemID,
			expectedType:   expectedItemType,
			expectedResult: true,
		},
		{
			name:  "Bad case (Custom ID)",
			input: expectedItemCustomID,
		},
		{
			name:  "Bad case (Effects custom ID)",
			input: expectedItemEffectsCustomID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualItemID, actualItemType, ok := ExtractItemRecipeCustomID(tt.input)
			assert.Equal(t, tt.expectedResult, ok)
			if ok {
				assert.Equal(t, tt.expectedID, actualItemID)
				assert.Equal(t, tt.expectedType, actualItemType)
			}
		})
	}
}

func TestIsBelongsToItem(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid custom ID",
			input:    expectedItemCustomID,
			expected: true,
		},
		{
			name:     "Valid effects custom ID",
			input:    expectedItemEffectsCustomID,
			expected: true,
		},
		{
			name:     "Valid recipe custom ID",
			input:    expectedItemRecipeCustomID,
			expected: true,
		},
		{
			name:  "Invalid custom ID",
			input: "/other/123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsBelongsToItem(tt.input))
		})
	}
}
