package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	ConfigCommandName           = "config"
	ConfigGetSubCommandName     = "get"
	ConfigAlmanaxSubCommandName = "almanax"
	ConfigRSSSubCommandName     = "rss"
	ConfigServerSubCommandName  = "server"
	ConfigTwitterSubCommandName = "twitter"

	ConfigServerOptionName         = "server"
	ConfigChannelOptionName        = "channel"
	ConfigFeedTypeOptionName       = "type"
	ConfigTwitterAccountOptionName = "account"
	ConfigEnabledOptionName        = "enabled"
)

//nolint:nolintlint,exhaustive,lll,dupl,funlen
func getConfigSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name: ConfigCommandName,
		Description: i18n.GetDefault("config.description",
			i18n.Vars{"game": constants.GetGame()}),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetManageServerPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("config.description",
			i18n.Vars{"game": constants.GetGame()}),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     ConfigGetSubCommandName,
				Description:              i18n.GetDefault("config.get.description"),
				NameLocalizations:        *i18n.GetLocalizations("config.get.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.get.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:                     ConfigAlmanaxSubCommandName,
				Description:              i18n.GetDefault("config.almanax.description"),
				NameLocalizations:        *i18n.GetLocalizations("config.almanax.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.almanax.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     ConfigEnabledOptionName,
						Description:              i18n.GetDefault("config.almanax.enabled.description"),
						NameLocalizations:        *i18n.GetLocalizations("config.almanax.enabled.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.almanax.enabled.description"),
						Type:                     discordgo.ApplicationCommandOptionBoolean,
						Required:                 true,
					},
					{
						Name:                     ConfigChannelOptionName,
						Description:              i18n.GetDefault("config.almanax.channel.description"),
						NameLocalizations:        *i18n.GetLocalizations("config.almanax.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.almanax.channel.description"),
						Type:                     discordgo.ApplicationCommandOptionChannel,
						Required:                 false,
					},
				},
			},
			{
				Name:                     ConfigRSSSubCommandName,
				Description:              i18n.GetDefault("config.rss.description"),
				NameLocalizations:        *i18n.GetLocalizations("config.rss.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.rss.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     ConfigEnabledOptionName,
						Description:              i18n.GetDefault("config.rss.enabled.description"),
						NameLocalizations:        *i18n.GetLocalizations("config.rss.enabled.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.rss.enabled.description"),
						Type:                     discordgo.ApplicationCommandOptionBoolean,
						Required:                 true,
					},
					{
						Name:                     ConfigFeedTypeOptionName,
						Description:              i18n.GetDefault("config.rss.feedtype.description"),
						NameLocalizations:        *i18n.GetLocalizations("config.rss.feedtype.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.rss.feedtype.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     ConfigChannelOptionName,
						Description:              i18n.GetDefault("config.rss.channel.description"),
						NameLocalizations:        *i18n.GetLocalizations("config.rss.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.rss.channel.description"),
						Type:                     discordgo.ApplicationCommandOptionChannel,
						Required:                 false,
					},
				},
			},
			{
				Name: ConfigServerSubCommandName,
				Description: i18n.GetDefault("config.server.description",
					i18n.Vars{"game": constants.GetGame()}),
				NameLocalizations: *i18n.GetLocalizations("config.server.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.server.description",
					i18n.Vars{"game": constants.GetGame()}),
				Type: discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name: ConfigServerOptionName,
						Description: i18n.GetDefault("config.server.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						NameLocalizations: *i18n.GetLocalizations("config.server.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.server.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						Type:         discordgo.ApplicationCommandOptionString,
						Required:     true,
						Autocomplete: true,
					},
					{
						Name: ConfigChannelOptionName,
						Description: i18n.GetDefault("config.server.channel.description",
							i18n.Vars{"game": constants.GetGame()}),
						NameLocalizations: *i18n.GetLocalizations("config.server.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.server.channel.description",
							i18n.Vars{"game": constants.GetGame()}),
						Type:     discordgo.ApplicationCommandOptionChannel,
						Required: false,
					},
				},
			},
			{
				Name: ConfigTwitterSubCommandName,
				Description: i18n.GetDefault("config.twitter.description",
					i18n.Vars{"game": constants.GetGame()}),
				NameLocalizations: *i18n.GetLocalizations("config.twitter.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.twitter.description",
					i18n.Vars{"game": constants.GetGame()}),
				Type: discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     ConfigEnabledOptionName,
						Description:              i18n.GetDefault("config.twitter.enabled.description"),
						NameLocalizations:        *i18n.GetLocalizations("config.twitter.enabled.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.twitter.enabled.description"),
						Type:                     discordgo.ApplicationCommandOptionBoolean,
						Required:                 true,
					},
					{
						Name:                     ConfigTwitterAccountOptionName,
						Description:              i18n.GetDefault("config.twitter.account.description"),
						NameLocalizations:        *i18n.GetLocalizations("config.twitter.account.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.twitter.account.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     ConfigChannelOptionName,
						Description:              i18n.GetDefault("config.twitter.channel.description"),
						NameLocalizations:        *i18n.GetLocalizations("config.twitter.channel.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("config.twitter.channel.description"),
						Type:                     discordgo.ApplicationCommandOptionChannel,
						Required:                 false,
					},
				},
			},
		},
	}
}
