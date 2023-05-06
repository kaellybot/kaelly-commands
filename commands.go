package commands

import "github.com/bwmarrin/discordgo"

type GetLocalChoicesFn func() []*discordgo.ApplicationCommandOptionChoice

func GetCommands(getLocalChoices GetLocalChoicesFn) []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		getAboutSlashCommand(),
		getAlignSlashCommand(),
		getAlignUserCommand(),
		getConfigSlashCommand(getLocalChoices),
		getJobSlashCommand(),
		getJobUserCommand(),
		getPosSlashCommand(),
	}
}
