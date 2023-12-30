package commands_test

import (
	"testing"

	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
)

const (
	expectedHelpCustomID = "/help"
)

func TestCraftHelpCustomID(t *testing.T) {
	actual := commands.CraftHelpCustomID()
	assert.Equal(t, expectedHelpCustomID, actual,
		"CraftHelpCustomID should return the expected Item custom ID")
}
