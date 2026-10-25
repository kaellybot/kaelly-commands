package commands

import "github.com/bwmarrin/discordgo"

func GetDofusCommands() []*discordgo.ApplicationCommand {
	g := getDofusGame()
	return []*discordgo.ApplicationCommand{
		getAboutSlashCommand(),
		getAlignSlashCommand(g),
		getAlignUserCommand(),
		getAlmanaxSlashCommand(),
		getConfigSlashCommand(g,
			getConfigGetSubCommand(),
			getConfigAlmanaxSubCommand(),
			getConfigRSSSubCommand(),
			getConfigServerSubCommand(g),
			getConfigTwitterSubCommand(g),
		),
		getHelpSlashCommand(),
		getItemSlashCommand(),
		getJobSlashCommand(g),
		getJobUserCommand(),
		getMapSlashCommand(),
		getSetSlashCommand(),
	}
}

func GetTouchCommands() []*discordgo.ApplicationCommand {
	g := getTouchGame()
	return []*discordgo.ApplicationCommand{
		getAboutSlashCommand(),
		getAlignSlashCommand(g),
		getAlignUserCommand(),
		getAlmanaxSlashCommand(),
		getConfigSlashCommand(g,
			getConfigGetSubCommand(),
			getConfigAlmanaxSubCommand(),
			getConfigRSSSubCommand(),
			getConfigServerSubCommand(g),
			getConfigTwitterSubCommand(g),
		),
		getHelpSlashCommand(),
		getJobSlashCommand(g),
		getJobUserCommand(),
	}
}

func GetRetroCommands() []*discordgo.ApplicationCommand {
	g := getRetroGame()
	return []*discordgo.ApplicationCommand{
		getAboutSlashCommand(),
		getAlignSlashCommand(g),
		getAlignUserCommand(),
		getConfigSlashCommand(g,
			getConfigGetSubCommand(),
			getConfigServerSubCommand(g),
		),
		getHelpSlashCommand(),
		getJobSlashCommand(g),
		getJobUserCommand(),
	}
}
