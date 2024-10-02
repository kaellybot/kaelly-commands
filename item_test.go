package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
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
	actual := commands.CraftItemCustomID(expectedItemType)
	assert.Equal(t, expectedItemCustomID, actual,
		"CraftItemCustomID should return the expected Item custom ID")
}

func TestCraftItemEffectsCustomID(t *testing.T) {
	actual := commands.CraftItemEffectsCustomID(expectedItemID, expectedItemType)
	assert.Equal(t, expectedItemEffectsCustomID, actual,
		"CraftItemEffectsCustomID should return the expected Effects custom ID")
}

func TestCraftItemRecipeCustomID(t *testing.T) {
	actual := commands.CraftItemRecipeCustomID(expectedItemID, expectedItemType)
	assert.Equal(t, expectedItemRecipeCustomID, actual,
		"CraftItemRecipeCustomID should return the expected Recipe custom ID")
}

func TestExtractItemCustomID(t *testing.T) {
	// Nominal case
	actualItemType, ok := commands.ExtractItemCustomID(expectedItemCustomID)
	assert.True(t, ok, "ExtractItemCustomID should indicate a successful extraction")
	assert.Equal(t, expectedItemType, actualItemType,
		"ExtractItemCustomID should return the expected Item type")

	// Bad case
	_, ok = commands.ExtractItemCustomID(expectedItemEffectsCustomID)
	assert.False(t, ok, "ExtractItemCustomID should indicate a failed extraction")
	_, ok = commands.ExtractItemCustomID(expectedItemRecipeCustomID)
	assert.False(t, ok, "ExtractItemCustomID should indicate a failed extraction")
}

func TestExtractItemEffectsCustomID(t *testing.T) {
	// Nominal case
	actualItemID, actualItemType, ok := commands.ExtractItemEffectsCustomID(expectedItemEffectsCustomID)
	assert.True(t, ok,
		"ExtractItemEffectsCustomID should indicate a successful extraction")
	assert.Equal(t, expectedItemID, actualItemID,
		"ExtractItemEffectsCustomID should return the expected Item ID")
	assert.Equal(t, expectedItemType, actualItemType,
		"ExtractItemEffectsCustomID should return the expected Item type")

	// Bad case
	_, _, ok = commands.ExtractItemEffectsCustomID(expectedItemCustomID)
	assert.False(t, ok, "ExtractItemEffectsCustomID should indicate a failed extraction")
	_, _, ok = commands.ExtractItemEffectsCustomID(expectedItemRecipeCustomID)
	assert.False(t, ok, "ExtractItemEffectsCustomID should indicate a failed extraction")
}

func TestExtractItemRecipeCustomID(t *testing.T) {
	// Nominal case
	actualItemID, actualItemType, ok := commands.ExtractItemRecipeCustomID(expectedItemRecipeCustomID)
	assert.True(t, ok,
		"ExtractItemRecipeCustomID should indicate a successful extraction")
	assert.Equal(t, expectedItemID, actualItemID,
		"ExtractItemRecipeCustomID should return the expected Item ID")
	assert.Equal(t, expectedItemType, actualItemType,
		"ExtractItemRecipeCustomID should return the expected Item type")

	// Bad case
	_, _, ok = commands.ExtractItemRecipeCustomID(expectedItemCustomID)
	assert.False(t, ok, "ExtractItemRecipeCustomID should indicate a failed extraction")
	_, _, ok = commands.ExtractItemRecipeCustomID(expectedItemEffectsCustomID)
	assert.False(t, ok, "ExtractItemRecipeCustomID should indicate a failed extraction")
}

func TestIsBelongsToItem(t *testing.T) {
	assert.True(t, commands.IsBelongsToItem(expectedItemCustomID),
		"IsBelongsToItem should return true for a valid Item custom ID")
	assert.True(t, commands.IsBelongsToItem(expectedItemEffectsCustomID),
		"IsBelongsToItem should return true for a valid Effects custom ID")
	assert.True(t, commands.IsBelongsToItem(expectedItemRecipeCustomID),
		"IsBelongsToItem should return true for a valid Recipe custom ID")
	assert.False(t, commands.IsBelongsToItem("/other/123"),
		"IsBelongsToItem should return false for a custom ID that doesn't belong to a Item")
}
