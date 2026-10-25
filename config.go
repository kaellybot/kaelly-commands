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

// notificationChannelTypes restricts the picker to what Discord accepts as the
// target of a followed announcement channel: a plain guild text channel. A
// thread holds no webhook of its own, and no other type can follow.
func notificationChannelTypes() []discordgo.ChannelType {
	return []discordgo.ChannelType{discordgo.ChannelTypeGuildText}
}

// serverChannelTypes restricts the picker to the channels a binding can ever be
// read back from: ones that host interactions, or that own the threads hosting
// them. A category hosts neither, so binding it would never apply.
func serverChannelTypes() []discordgo.ChannelType {
	return []discordgo.ChannelType{
		discordgo.ChannelTypeGuildText,
		discordgo.ChannelTypeGuildVoice,
		discordgo.ChannelTypeGuildNews,
		discordgo.ChannelTypeGuildForum,
		discordgo.ChannelTypeGuildMedia,
	}
}

// getConfigSlashCommand builds /config with only the given sub-commands: each
// game configures a different set of features.
//
//nolint:nolintlint,exhaustive,lll,dupl
func getConfigSlashCommand(g game, subCommands ...*discordgo.ApplicationCommandOption,
) *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name: ConfigCommandName,
		Description: i18n.GetDefault("config.description",
			i18n.Vars{"game": g.name}),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetManageServerPermission(),
		Contexts:                 constants.GetContexts(),
		DescriptionLocalizations: i18n.GetLocalizations("config.description",
			i18n.Vars{"game": g.name}),
		Options: subCommands,
	}
}

//nolint:nolintlint,exhaustive,lll,dupl
func getConfigGetSubCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:                     ConfigGetSubCommandName,
		Description:              i18n.GetDefault("config.get.description"),
		NameLocalizations:        *i18n.GetLocalizations("config.get.name"),
		DescriptionLocalizations: *i18n.GetLocalizations("config.get.description"),
		Type:                     discordgo.ApplicationCommandOptionSubCommand,
	}
}

//nolint:nolintlint,exhaustive,lll,dupl
func getConfigAlmanaxSubCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
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
				ChannelTypes:             notificationChannelTypes(),
				Required:                 false,
			},
		},
	}
}

//nolint:nolintlint,exhaustive,lll,dupl
func getConfigRSSSubCommand() *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
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
				ChannelTypes:             notificationChannelTypes(),
				Required:                 false,
			},
		},
	}
}

//nolint:nolintlint,exhaustive,lll,dupl
func getConfigServerSubCommand(g game) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name: ConfigServerSubCommandName,
		Description: i18n.GetDefault("config.server.description",
			i18n.Vars{"game": g.name}),
		NameLocalizations: *i18n.GetLocalizations("config.server.name"),
		DescriptionLocalizations: *i18n.GetLocalizations("config.server.description",
			i18n.Vars{"game": g.name}),
		Type: discordgo.ApplicationCommandOptionSubCommand,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: ConfigServerOptionName,
				Description: i18n.GetDefault("config.server.server.description",
					i18n.Vars{"game": g.name}),
				NameLocalizations: *i18n.GetLocalizations("config.server.server.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.server.server.description",
					i18n.Vars{"game": g.name}),
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
			},
			{
				Name: ConfigChannelOptionName,
				Description: i18n.GetDefault("config.server.channel.description",
					i18n.Vars{"game": g.name}),
				NameLocalizations: *i18n.GetLocalizations("config.server.channel.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("config.server.channel.description",
					i18n.Vars{"game": g.name}),
				Type:         discordgo.ApplicationCommandOptionChannel,
				ChannelTypes: serverChannelTypes(),
				Required:     false,
			},
		},
	}
}

//nolint:nolintlint,exhaustive,lll,dupl
func getConfigTwitterSubCommand(g game) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name: ConfigTwitterSubCommandName,
		Description: i18n.GetDefault("config.twitter.description",
			i18n.Vars{"game": g.name}),
		NameLocalizations: *i18n.GetLocalizations("config.twitter.name"),
		DescriptionLocalizations: *i18n.GetLocalizations("config.twitter.description",
			i18n.Vars{"game": g.name}),
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
				ChannelTypes:             notificationChannelTypes(),
				Required:                 false,
			},
		},
	}
}
