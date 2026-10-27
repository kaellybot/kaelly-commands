package commands_test

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	commands "github.com/kaellybot/kaelly-commands"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// configChannelOption returns the channel option of the given /config
// sub-command.
func configChannelOption(t *testing.T, cmds []*discordgo.ApplicationCommand,
	subCommandName string) *discordgo.ApplicationCommandOption {
	t.Helper()
	config := findCommand(t, cmds, commands.ConfigCommandName)
	subCommand := findOption(t, config.Options, subCommandName)
	return findOption(t, subCommand.Options, commands.ConfigChannelOptionName)
}

// Notifications are set up by following an announcement channel, and Discord
// only lets a plain text channel follow one: the picker must not offer the
// rest, otherwise the bot can only refuse after the fact.
func TestConfigNotificationChannelOptionsOnlyOfferTextChannels(t *testing.T) {
	for game, cmds := range gameCommands() {
		subCommandNames := []string{commands.ConfigRSSSubCommandName, commands.ConfigTwitterSubCommandName}
		if game != "DOFUS Retro" {
			subCommandNames = append(subCommandNames, commands.ConfigAlmanaxSubCommandName)
		}

		for _, subCommandName := range subCommandNames {
			t.Run(game+"/"+subCommandName, func(t *testing.T) {
				option := configChannelOption(t, cmds, subCommandName)

				assert.Equal(t, []discordgo.ChannelType{discordgo.ChannelTypeGuildText}, option.ChannelTypes)
			})
		}
	}
}

// A server binding is read back from the channel an interaction comes from, or
// from the parent of the thread hosting it: a category is neither.
func TestConfigServerChannelOptionOffersChannelsHostingInteractions(t *testing.T) {
	for game, cmds := range gameCommands() {
		t.Run(game, func(t *testing.T) {
			option := configChannelOption(t, cmds, commands.ConfigServerSubCommandName)

			require.NotEmpty(t, option.ChannelTypes)
			assert.Contains(t, option.ChannelTypes, discordgo.ChannelTypeGuildText)
			assert.Contains(t, option.ChannelTypes, discordgo.ChannelTypeGuildForum,
				"a forum owns the threads its posts live in")
			for _, refused := range []discordgo.ChannelType{
				discordgo.ChannelTypeGuildCategory,
				discordgo.ChannelTypeGuildPublicThread,
				discordgo.ChannelTypeGuildPrivateThread,
				discordgo.ChannelTypeDM,
			} {
				assert.NotContains(t, option.ChannelTypes, refused)
			}
		})
	}
}
