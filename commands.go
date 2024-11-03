package commands

import "github.com/bwmarrin/discordgo"

func GetCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		getAboutSlashCommand(),
		getAlignSlashCommand(),
		getAlignUserCommand(),
		getAlmanaxSlashCommand(),
		getConfigSlashCommand(),
		getHelpSlashCommand(),
		getItemSlashCommand(),
		getJobSlashCommand(),
		getJobUserCommand(),
		getMapSlashCommand(),
		getPosSlashCommand(),
		getSetSlashCommand(),
	}
}
