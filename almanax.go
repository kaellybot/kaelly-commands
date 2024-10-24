package commands

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	"github.com/kaellybot/kaelly-commands/utils/regex"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	AlmanaxCommandName             = "almanax"
	AlmanaxDaySubCommandName       = "day"
	AlmanaxEffectsSubCommandName   = "effects"
	AlmanaxResourcesSubCommandName = "resources"

	AlmanaxDateOptionName     = "date"
	AlmanaxEffectOptionName   = "effect"
	AlmanaxDurationOptionName = "duration"

	AlmanaxDurationMinimumValue = 1.0
	AlmanaxDurationDefaultValue = 7.0
	AlmanaxDurationMaximumValue = 30.0

	almanaxDayCustomIDGroups = 2
)

var (
	AlmanaxDayCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/day/(\\d+)$", AlmanaxCommandName))
)

//nolint:nolintlint,exhaustive,lll,dupl,funlen
func getAlmanaxSlashCommand() *discordgo.ApplicationCommand {
	almanaxDurationMinimumValue := AlmanaxDurationMinimumValue
	return &discordgo.ApplicationCommand{
		Name:                     AlmanaxCommandName,
		Description:              "almanax.description",
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("almanax.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     AlmanaxDaySubCommandName,
				Description:              "almanax.day.description",
				NameLocalizations:        *i18n.GetLocalizations("almanax.day.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("almanax.day.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     AlmanaxDateOptionName,
						Description:              "almanax.day.date.description",
						NameLocalizations:        *i18n.GetLocalizations("almanax.day.date.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("almanax.day.date.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 false,
					},
				},
			},
			{
				Name:                     AlmanaxResourcesSubCommandName,
				Description:              "almanax.resources.description",
				NameLocalizations:        *i18n.GetLocalizations("almanax.resources.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("almanax.resources.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     AlmanaxDurationOptionName,
						Description:              "almanax.resources.duration.description",
						NameLocalizations:        *i18n.GetLocalizations("almanax.resources.duration.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("almanax.resources.duration.description"),
						Type:                     discordgo.ApplicationCommandOptionInteger,
						Required:                 false,
						MinValue:                 &almanaxDurationMinimumValue,
						MaxValue:                 AlmanaxDurationMaximumValue,
					},
				},
			},
			{
				Name:                     AlmanaxEffectsSubCommandName,
				Description:              "almanax.effects.description",
				NameLocalizations:        *i18n.GetLocalizations("almanax.effects.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("almanax.effects.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     AlmanaxEffectOptionName,
						Description:              "almanax.effects.effect.description",
						NameLocalizations:        *i18n.GetLocalizations("almanax.effects.effect.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("almanax.effects.effect.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
				},
			},
		},
	}
}

func CraftAlmanaxDayCustomID(date time.Time) string {
	return fmt.Sprintf("/%s/day/%v", AlmanaxCommandName, date.Unix())
}

func ExtractAlmanaxDayCustomID(customID string) (*time.Time, bool) {
	if groups, ok := regex.ExtractCustomID(customID, AlmanaxDayCustomID,
		almanaxDayCustomIDGroups); ok {
		seconds, err := strconv.ParseInt(groups[1], 10, 64)
		if err != nil {
			return nil, false
		}

		day := time.Unix(seconds, 0).UTC()
		return &day, true
	}

	return nil, false
}

func IsBelongsToAlmanax(customID string) bool {
	return regex.IsBelongTo(customID, AlmanaxDayCustomID)
}
