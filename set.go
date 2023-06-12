package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	SetCommandName = "set"

	SetQueryOptionName = "query"
)

//nolint:nolintlint,exhaustive,lll,dupl,funlen
func getSetSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     SetCommandName,
		Description:              "set.description",
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("set.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     SetQueryOptionName,
				Description:              "set.query.description",
				NameLocalizations:        *i18n.GetLocalizations("set.query.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("set.query.description"),
				Type:                     discordgo.ApplicationCommandOptionString,
				Required:                 true,
				Autocomplete:             true,
			},
		},
	}
}
