package commands

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	"github.com/kaellybot/kaelly-commands/utils/regex"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	AlignSlashCommandName  = "align"
	AlignUserCommandName   = "Alignments"
	AlignGetSubCommandName = "get"
	AlignSetSubCommandName = "set"

	AlignCityOptionName   = "city"
	AlignOrderOptionName  = "order"
	AlignLevelOptionName  = "level"
	AlignServerOptionName = "server"

	AlignmentMinLevel = 0
	AlignmentMaxLevel = 100

	alignBookPageCustomIDGroups        = 5
	alignBookCitySelectCustomIDGroups  = 3
	alignBookOrderSelectCustomIDGroups = 3

	AlignAllValues = "_"
)

var (
	AlignBookPageCustomID = regexp.MustCompile(fmt.
				Sprintf("^/books/%s\\?city=([a-z_]+)&order=([a-z_]+)&server=([a-z_]+)&page=(\\d+)$", AlignSlashCommandName))
	AlignBookCityCustomID = regexp.MustCompile(fmt.
				Sprintf("^/books/%s\\?order=([a-z_]+)&server=([a-z_]+)$", AlignSlashCommandName))
	AlignBookOrderCustomID = regexp.MustCompile(fmt.
				Sprintf("^/books/%s\\?city=([a-z_]+)&server=([a-z_]+)$", AlignSlashCommandName))
)

//nolint:nolintlint,exhaustive,lll,dupl
func getAlignSlashCommand() *discordgo.ApplicationCommand {
	var minLevel float64 = AlignmentMinLevel
	return &discordgo.ApplicationCommand{
		Name:                     AlignSlashCommandName,
		Description:              i18n.GetDefault("align.description"),
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("align.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     AlignGetSubCommandName,
				Description:              i18n.GetDefault("align.get.description"),
				NameLocalizations:        *i18n.GetLocalizations("align.get.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("align.get.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     AlignCityOptionName,
						Description:              i18n.GetDefault("align.get.city.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.get.city.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.get.city.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 false,
						Autocomplete:             true,
					},
					{
						Name:                     AlignOrderOptionName,
						Description:              i18n.GetDefault("align.get.order.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.get.order.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.get.order.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 false,
						Autocomplete:             true,
					},
					{
						Name: AlignServerOptionName,
						Description: i18n.GetDefault("align.get.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						NameLocalizations: *i18n.GetLocalizations("align.get.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.get.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						Type:         discordgo.ApplicationCommandOptionString,
						Required:     false,
						Autocomplete: true,
					},
				},
			},
			{
				Name:                     AlignSetSubCommandName,
				Description:              i18n.GetDefault("align.set.description"),
				NameLocalizations:        *i18n.GetLocalizations("align.set.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("align.set.description"),
				Type:                     discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:                     AlignCityOptionName,
						Description:              i18n.GetDefault("align.set.city.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.set.city.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.city.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     AlignOrderOptionName,
						Description:              i18n.GetDefault("align.set.order.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.set.order.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.order.description"),
						Type:                     discordgo.ApplicationCommandOptionString,
						Required:                 true,
						Autocomplete:             true,
					},
					{
						Name:                     AlignLevelOptionName,
						Description:              i18n.GetDefault("align.set.level.description"),
						NameLocalizations:        *i18n.GetLocalizations("align.set.level.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.level.description"),
						Type:                     discordgo.ApplicationCommandOptionInteger,
						Required:                 true,
						MinValue:                 &minLevel,
						MaxValue:                 AlignmentMaxLevel,
					},
					{
						Name: AlignServerOptionName,
						Description: i18n.GetDefault("align.set.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						NameLocalizations: *i18n.GetLocalizations("align.set.server.name"),
						DescriptionLocalizations: *i18n.GetLocalizations("align.set.server.description",
							i18n.Vars{"game": constants.GetGame()}),
						Type:         discordgo.ApplicationCommandOptionString,
						Required:     false,
						Autocomplete: true,
					},
				},
			},
		},
	}
}

//nolint:nolintlint,exhaustive,lll,dupl
func getAlignUserCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     AlignUserCommandName,
		Type:                     discordgo.UserApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
	}
}

func CraftAlignBookPageCustomID(cityID, orderID, serverID string, page int) string {
	if len(cityID) == 0 {
		cityID = AlignAllValues
	}

	if len(orderID) == 0 {
		orderID = AlignAllValues
	}

	return fmt.Sprintf("/books/%s?city=%v&order=%v&server=%v&page=%v",
		AlignSlashCommandName, cityID, orderID, serverID, page)
}

func ExtractAlignBookPageCustomID(customID string,
) (string, string, string, int, bool) {
	if groups, ok := regex.ExtractCustomID(customID, AlignBookPageCustomID,
		alignBookPageCustomIDGroups); ok {
		cityID := groups[1]
		orderID := groups[2]
		serverID := groups[3]

		if cityID == AlignAllValues {
			cityID = ""
		}

		if orderID == AlignAllValues {
			orderID = ""
		}

		page, errConv := strconv.Atoi(groups[4])
		if errConv != nil {
			return "", "", "", -1, false
		}

		return cityID, orderID, serverID, page, true
	}

	return "", "", "", -1, false
}

func CraftAlignBookCityCustomID(orderID, serverID string) string {
	if len(orderID) == 0 {
		orderID = AlignAllValues
	}

	return fmt.Sprintf("/books/%s?order=%v&server=%v",
		AlignSlashCommandName, orderID, serverID)
}

func ExtractAlignBookCityCustomID(customID string,
) (string, string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, AlignBookCityCustomID,
		alignBookCitySelectCustomIDGroups); ok {
		orderID := groups[1]
		serverID := groups[2]

		if orderID == AlignAllValues {
			orderID = ""
		}

		return orderID, serverID, true
	}

	return "", "", false
}

func ExtractAlignBookOrderCustomID(customID string,
) (string, string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, AlignBookOrderCustomID,
		alignBookOrderSelectCustomIDGroups); ok {
		cityID := groups[1]
		serverID := groups[2]

		if cityID == AlignAllValues {
			cityID = ""
		}

		return cityID, serverID, true
	}

	return "", "", false
}

func CraftAlignBookOrderCustomID(cityID, serverID string) string {
	if len(cityID) == 0 {
		cityID = AlignAllValues
	}

	return fmt.Sprintf("/books/%s?city=%v&server=%v",
		AlignSlashCommandName, cityID, serverID)
}

func IsBelongsToAlign(customID string) bool {
	return regex.IsBelongTo(customID, AlignBookPageCustomID,
		AlignBookCityCustomID, AlignBookOrderCustomID)
}
