package commands

import "github.com/bwmarrin/discordgo"

func GetCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		getAboutSlashCommand(),
		getAlignSlashCommand(),
		getAlignUserCommand(),
		getConfigSlashCommand(),
		getJobSlashCommand(),
		getJobUserCommand(),
		getPosSlashCommand(),
	}
}
