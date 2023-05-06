package commands

import (
	"github.com/bwmarrin/discordgo"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	JobSlashCommandName  = "job"
	JobUserCommandName   = "Jobs"
	JobGetSubCommandName = "get"
	JobSetSubCommandName = "set"

	JobJobOptionName    = "job"
	JobLevelOptionName  = "level"
	JobServerOptionName = "server"

	JobMinLevel = 0
	JobMaxLevel = 200
)

//nolint:nolintlint,exhaustive,lll,dupl
func getJobSlashCommand() *discordgo.ApplicationCommand {
	var minLevel float64 = JobMinLevel
	return &discordgo.ApplicationCommand{
		Name:                     JobSlashCommandName,
		Description:              i18n.Get(DefaultLocale, "job.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: GetDefaultPermission(),
		DMPermission:             GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("job.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     JobGetSubCommandName,
				Description:              i18n.Get(DefaultLocale, "job.get.description"),
				NameLocalizations:        *i18n.GetLocalizations("job.get.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("job.get.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     JobJobOptionName,
						Description:              i18n.Get(DefaultLocale, "job.get.job.description"),
						NameLocalizations:        *i18n.GetLocalizations("job.get.job.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.get.job.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     JobServerOptionName,
						Description:              i18n.Get(DefaultLocale, "job.get.server.description", i18n.Vars{"game": GetGame()}),
						NameLocalizations:        *i18n.GetLocalizations("job.get.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.get.server.description", i18n.Vars{"game": GetGame()}),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 false,
						Autocomplete:             true,
					},
				},
			},
			{
				Name:                     JobSetSubCommandName,
				Description:              i18n.Get(DefaultLocale, "job.set.description"),
				NameLocalizations:        *i18n.GetLocalizations("job.set.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("job.set.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     JobJobOptionName,
						Description:              i18n.Get(DefaultLocale, "job.set.job.description"),
						NameLocalizations:        *i18n.GetLocalizations("job.set.job.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.set.job.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     JobLevelOptionName,
						Description:              i18n.Get(DefaultLocale, "job.set.level.description"),
						NameLocalizations:        *i18n.GetLocalizations("job.set.level.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.set.level.description"),
						Type:                     discordgo.ApplicationCommandOptionInteger,
						Required:                 true,
						MinValue:                 &minLevel,
						MaxValue:                 JobMaxLevel,
					},
					{
						Name:                     JobServerOptionName,
						Description:              i18n.Get(DefaultLocale, "job.set.server.description", i18n.Vars{"game": GetGame()}),
						NameLocalizations:        *i18n.GetLocalizations("job.set.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("job.set.server.description", i18n.Vars{"game": GetGame()}),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 false,
						Autocomplete:             true,
					},
				},
			},
		},
	}
}

//nolint:nolintlint,exhaustive,lll,dupl
func getJobUserCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     JobUserCommandName,
		Type:                     discordgo.UserApplicationCommand,
		DefaultMemberPermissions: GetDefaultPermission(),
		DMPermission:             GetDMPermission(),
	}
}
