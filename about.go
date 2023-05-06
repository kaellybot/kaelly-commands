package commands

import (
	"github.com/bwmarrin/discordgo"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	AboutCommandName = "about"
)

//nolint:nolintlint,exhaustive,lll,dupl
func getAboutSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     AboutCommandName,
		Description:              i18n.Get(DefaultLocale, "about.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: GetDefaultPermission(),
		DMPermission:             GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("about.description"),
	}
}
