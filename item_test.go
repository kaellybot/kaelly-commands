package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedItemID             = "123"
	expectedItemCustomID       = "/item/123"
	expectedItemRecipeCustomID = "/item/123/recipe"
)

func TestCraftItemCustomID(t *testing.T) {
	actual := CraftItemCustomID(expectedItemID)
	assert.Equal(t, expectedItemCustomID, actual, "CraftItemCustomID should return the expected Item custom ID")
}

func TestCraftItemRecipeCustomID(t *testing.T) {
	actual := CraftItemRecipeCustomID(expectedItemID)
	assert.Equal(t, expectedItemRecipeCustomID, actual, "CraftItemRecipeCustomID should return the expected Recipe custom ID")
}

func TestExtractItemCustomID(t *testing.T) {
	// Nominal case
	actualItemID, ok := ExtractItemCustomID(expectedItemCustomID)
	assert.True(t, ok, "ExtractItemCustomID should indicate a successful extraction")
	assert.Equal(t, expectedItemID, actualItemID, "ExtractItemCustomID should return the expected Item ID")

	// Bad case
	_, ok = ExtractItemCustomID(expectedItemRecipeCustomID)
	assert.False(t, ok, "expectedItemCustomID should indicate a failed extraction")
}

func TestExtractItemRecipeCustomID(t *testing.T) {
	// Nominal case
	actualItemID, ok := ExtractItemRecipeCustomID(expectedItemRecipeCustomID)
	assert.True(t, ok, "ExtractItemRecipeCustomID should indicate a successful extraction")
	assert.Equal(t, expectedItemID, actualItemID, "ExtractItemRecipeCustomID should return the expected Item ID")

	// Bad case
	_, ok = ExtractItemRecipeCustomID(expectedItemCustomID)
	assert.False(t, ok, "ExtractItemRecipeCustomID should indicate a failed extraction")
}

func TestIsBelongsToItem(t *testing.T) {
	assert.True(t, IsBelongsToItem(expectedItemCustomID), "IsBelongsToItem should return true for a valid Item custom ID")
	assert.True(t, IsBelongsToItem(expectedItemRecipeCustomID), "IsBelongsToItem should return true for a valid Recipe custom ID")
	assert.False(t, IsBelongsToItem("/other/123"), "IsBelongsToItem should return false for a custom ID that doesn't belong to a Item")
}
