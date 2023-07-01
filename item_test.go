package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedItemID             = "123"
	expectedItemCustomID       = "/item"
	expectedItemRecipeCustomID = "/item/123/recipe"
)

func TestCraftItemCustomID(t *testing.T) {
	actual := commands.CraftItemCustomID()
	assert.Equal(t, expectedItemCustomID, actual,
		"CraftItemCustomID should return the expected Item custom ID")
}

func TestCraftItemRecipeCustomID(t *testing.T) {
	actual := commands.CraftItemRecipeCustomID(expectedItemID)
	assert.Equal(t, expectedItemRecipeCustomID, actual,
		"CraftItemRecipeCustomID should return the expected Recipe custom ID")
}

func TestExtractItemRecipeCustomID(t *testing.T) {
	// Nominal case
	actualItemID, ok := commands.ExtractItemRecipeCustomID(expectedItemRecipeCustomID)
	assert.True(t, ok,
		"ExtractItemRecipeCustomID should indicate a successful extraction")
	assert.Equal(t, expectedItemID, actualItemID,
		"ExtractItemRecipeCustomID should return the expected Item ID")

	// Bad case
	_, ok = commands.ExtractItemRecipeCustomID(expectedItemCustomID)
	assert.False(t, ok, "ExtractItemRecipeCustomID should indicate a failed extraction")
}

func TestIsBelongsToItem(t *testing.T) {
	assert.True(t, commands.IsBelongsToItem(expectedItemCustomID),
		"IsBelongsToItem should return true for a valid Item custom ID")
	assert.True(t, commands.IsBelongsToItem(expectedItemRecipeCustomID),
		"IsBelongsToItem should return true for a valid Recipe custom ID")
	assert.False(t, commands.IsBelongsToItem("/other/123"),
		"IsBelongsToItem should return false for a custom ID that doesn't belong to a Item")
}
