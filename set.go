// nolint:nolintlint,dupl
package commands

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	"github.com/kaellybot/kaelly-commands/utils/regex"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	SetCommandName = "set"

	SetQueryOptionName = "query"

	setCustomIDGroups      = 2
	setBonusCustomIDGroups = 2
)

var (
	setCustomID      = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)$", SetCommandName))
	setBonusCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/bonuses$", SetCommandName))
)

//nolint:exhaustive,lll,funlen
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

func CraftSetCustomID(setID string) string {
	return fmt.Sprintf("/%s/%s", SetCommandName, setID)
}

func CraftSetBonusCustomID(setID string) string {
	return fmt.Sprintf("/%s/%s/bonuses", SetCommandName, setID)
}

func ExtractSetCustomID(customID string) (string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, setCustomID,
		setCustomIDGroups); ok {
		return groups[1], true
	}

	return "", false
}

func ExtractSetBonusCustomID(customID string) (string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, setBonusCustomID,
		setBonusCustomIDGroups); ok {
		return groups[1], true
	}

	return "", false
}

func IsBelongsToSet(customID string) bool {
	return regex.IsBelongTo(customID, setCustomID, setBonusCustomID)
}
