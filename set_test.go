package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedSetID            = "123"
	expectedBonus            = 7
	expectedSetCustomID      = "/set/123"
	expectedSetBonusCustomID = "/set/123/bonuses/7"
)

func TestCraftSetCustomID(t *testing.T) {
	actual := commands.CraftSetCustomID(expectedSetID)
	assert.Equal(t, expectedSetCustomID, actual,
		"CraftSetCustomID should return the expected set custom ID")
}

func TestCraftSetBonusCustomID(t *testing.T) {
	actual := commands.CraftSetBonusCustomID(expectedSetID, expectedBonus)
	assert.Equal(t, expectedSetBonusCustomID, actual,
		"CraftSetBonusCustomID should return the expected bonus custom ID")
}

func TestExtractSetCustomID(t *testing.T) {
	// Nominal case
	actualSetID, ok := commands.ExtractSetCustomID(expectedSetCustomID)
	assert.True(t, ok,
		"ExtractSetCustomID should indicate a successful extraction")
	assert.Equal(t, expectedSetID, actualSetID,
		"ExtractSetCustomID should return the expected set ID")

	// Bad case
	_, ok = commands.ExtractSetCustomID(expectedSetBonusCustomID)
	assert.False(t, ok, "expectedSetCustomID should indicate a failed extraction")
}

func TestExtractSetBonusCustomID(t *testing.T) {
	// Nominal case
	actualSetID, actualItemNb, ok := commands.ExtractSetBonusCustomID(expectedSetBonusCustomID)
	assert.True(t, ok,
		"ExtractSetBonusCustomID should indicate a successful extraction")
	assert.Equal(t, expectedSetID, actualSetID,
		"ExtractSetBonusCustomID should return the expected set ID")
	assert.Equal(t, expectedBonus, actualItemNb,
		"ExtractSetBonusCustomID should return the expected set ID")

	// Bad case
	_, _, ok = commands.ExtractSetBonusCustomID(expectedSetCustomID)
	assert.False(t, ok, "ExtractSetBonusCustomID should indicate a failed extraction")
}

func TestIsBelongsToSet(t *testing.T) {
	assert.True(t, commands.IsBelongsToSet(expectedSetCustomID),
		"IsBelongsToSet should return true for a valid set custom ID")
	assert.True(t, commands.IsBelongsToSet(expectedSetBonusCustomID),
		"IsBelongsToSet should return true for a valid bonus custom ID")
	assert.False(t, commands.IsBelongsToSet("/other/123"),
		"IsBelongsToSet should return false for a custom ID that doesn't belong to a set")
}
