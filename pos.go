package commands

import (
	"github.com/bwmarrin/discordgo"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	PosCommandName         = "pos"
	PosDimensionOptionName = "dimension"
	PosServerOptionName    = "server"
)

//nolint:nolintlint,exhaustive,lll,dupl
func getPosSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     PosCommandName,
		Description:              i18n.Get(DefaultLocale, "pos.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: GetDefaultPermission(),
		DMPermission:             GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("pos.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     PosDimensionOptionName,
				Description:              i18n.Get(DefaultLocale, "pos.dimension.description"),
				NameLocalizations:        *i18n.GetLocalizations("pos.dimension.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("pos.dimension.description"),
				Type:                     discordgo.ApplicationCommandOptionString,
				Required:                 false,
				Autocomplete:             true,
			},
			{
				Name:                     PosServerOptionName,
				Description:              i18n.Get(DefaultLocale, "pos.server.description", i18n.Vars{"game": GetGame()}),
				NameLocalizations:        *i18n.GetLocalizations("pos.server.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("pos.server.description", i18n.Vars{"game": GetGame()}),
				Type:                     discordgo.ApplicationCommandOptionString,
				Required:                 false,
				Autocomplete:             true,
			},
		},
	}
}
