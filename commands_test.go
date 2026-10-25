package commands_test

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	commands "github.com/kaellybot/kaelly-commands"
	i18n "github.com/kaysoro/discordgo-i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// gameCommands lists the command set of every game.
func gameCommands() map[string][]*discordgo.ApplicationCommand {
	return map[string][]*discordgo.ApplicationCommand{
		"DOFUS":       commands.GetDofusCommands(),
		"DOFUS Touch": commands.GetTouchCommands(),
		"DOFUS Retro": commands.GetRetroCommands(),
	}
}

func commandNames(cmds []*discordgo.ApplicationCommand) []string {
	names := make([]string, 0, len(cmds))
	for _, command := range cmds {
		names = append(names, command.Name)
	}
	return names
}

func findCommand(t *testing.T, cmds []*discordgo.ApplicationCommand, name string) *discordgo.ApplicationCommand {
	t.Helper()
	for _, command := range cmds {
		if command.Name == name {
			return command
		}
	}

	t.Fatalf("no /%s command found", name)
	return nil
}

func findOption(t *testing.T, options []*discordgo.ApplicationCommandOption,
	name string) *discordgo.ApplicationCommandOption {
	t.Helper()
	for _, option := range options {
		if option.Name == name {
			return option
		}
	}

	t.Fatalf("no %s option found", name)
	return nil
}

func subCommandNames(command *discordgo.ApplicationCommand) []string {
	names := make([]string, 0, len(command.Options))
	for _, option := range command.Options {
		names = append(names, option.Name)
	}
	return names
}

func TestEachGameRegistersItsOwnCommands(t *testing.T) {
	assert.Equal(t, []string{
		commands.AboutCommandName, commands.AlignSlashCommandName, commands.AlignUserCommandName,
		commands.AlmanaxCommandName, commands.ConfigCommandName, commands.HelpCommandName,
		commands.ItemCommandName, commands.JobSlashCommandName, commands.JobUserCommandName,
		commands.MapCommandName, commands.SetCommandName,
	}, commandNames(commands.GetDofusCommands()))
	assert.Equal(t, []string{
		commands.AboutCommandName, commands.AlignSlashCommandName, commands.AlignUserCommandName,
		commands.AlmanaxCommandName, commands.ConfigCommandName, commands.HelpCommandName,
		commands.JobSlashCommandName, commands.JobUserCommandName,
	}, commandNames(commands.GetTouchCommands()))
	assert.Equal(t, []string{
		commands.AboutCommandName, commands.AlignSlashCommandName, commands.AlignUserCommandName,
		commands.ConfigCommandName, commands.HelpCommandName,
		commands.JobSlashCommandName, commands.JobUserCommandName,
	}, commandNames(commands.GetRetroCommands()))
}

func TestEachGameConfiguresItsOwnFeatures(t *testing.T) {
	notifications := []string{
		commands.ConfigGetSubCommandName, commands.ConfigAlmanaxSubCommandName, commands.ConfigRSSSubCommandName,
		commands.ConfigServerSubCommandName, commands.ConfigTwitterSubCommandName,
	}
	assert.Equal(t, notifications,
		subCommandNames(findCommand(t, commands.GetDofusCommands(), commands.ConfigCommandName)))
	assert.Equal(t, notifications,
		subCommandNames(findCommand(t, commands.GetTouchCommands(), commands.ConfigCommandName)))
	assert.Equal(t, []string{commands.ConfigGetSubCommandName, commands.ConfigServerSubCommandName},
		subCommandNames(findCommand(t, commands.GetRetroCommands(), commands.ConfigCommandName)))
}

func TestLevelsAreCappedPerGame(t *testing.T) {
	expected := map[string]struct{ alignment, job float64 }{
		"DOFUS":       {commands.AlignmentMaxLevel, commands.JobMaxLevel},
		"DOFUS Touch": {commands.AlignmentTouchMaxLevel, commands.JobTouchMaxLevel},
		"DOFUS Retro": {commands.AlignmentRetroMaxLevel, commands.JobRetroMaxLevel},
	}

	for game, cmds := range gameCommands() {
		t.Run(game, func(t *testing.T) {
			alignSet := findOption(t, findCommand(t, cmds, commands.AlignSlashCommandName).Options,
				commands.AlignSetSubCommandName)
			alignLevel := findOption(t, alignSet.Options, commands.AlignLevelOptionName)
			require.NotNil(t, alignLevel.MinValue)
			assert.InDelta(t, commands.AlignmentMinLevel, *alignLevel.MinValue, 0)
			assert.InDelta(t, expected[game].alignment, alignLevel.MaxValue, 0)

			jobSet := findOption(t, findCommand(t, cmds, commands.JobSlashCommandName).Options,
				commands.JobSetSubCommandName)
			jobLevel := findOption(t, jobSet.Options, commands.JobLevelOptionName)
			require.NotNil(t, jobLevel.MinValue)
			assert.InDelta(t, commands.JobMinLevel, *jobLevel.MinValue, 0)
			assert.InDelta(t, expected[game].job, jobLevel.MaxValue, 0)
		})
	}
}

// Server descriptions name the game, so that players know which servers the
// autocompletion offers.
func TestDescriptionsNameTheirGame(t *testing.T) {
	require.NoError(t, i18n.LoadBundleContent(discordgo.EnglishGB, map[string]any{
		"config": map[string]any{
			"server": map[string]any{"description": "Links the guild to a {{ .game }} server"},
		},
	}))
	i18n.SetDefault(discordgo.EnglishGB)

	for game, cmds := range gameCommands() {
		t.Run(game, func(t *testing.T) {
			server := findOption(t, findCommand(t, cmds, commands.ConfigCommandName).Options,
				commands.ConfigServerSubCommandName)

			assert.Equal(t, "Links the guild to a "+game+" server", server.Description)
			assert.Equal(t, "Links the guild to a "+game+" server",
				server.DescriptionLocalizations[discordgo.EnglishGB])
		})
	}
}
