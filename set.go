package commands

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	SetCommandName = "set"

	SetQueryOptionName = "query"

	setCustomIDGroups      = 2
	setBonusCustomIDGroups = 3
)

var (
	setCustomID      = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)$", SetCommandName))
	setBonusCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/bonuses/(\\d+)$", SetCommandName))
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

func CraftSetCustomID(setID string) string {
	return fmt.Sprintf("/%s/%s", SetCommandName, setID)
}

func CraftSetBonusCustomID(setID string, itemNumber int) string {
	return fmt.Sprintf("/%s/%s/bonuses/%d", SetCommandName, setID, itemNumber)
}

func ExtractSetCustomID(customID string) (string, bool) {
	if setCustomID.MatchString(customID) {
		groups := setCustomID.FindStringSubmatch(customID)
		if len(groups) == setCustomIDGroups {
			return groups[1], true
		}
	}

	return "", false
}

func ExtractSetBonusCustomID(customID string) (string, int, bool) {
	if setBonusCustomID.MatchString(customID) {
		groups := setBonusCustomID.FindStringSubmatch(customID)
		if len(groups) == setBonusCustomIDGroups {
			itemNumber, err := strconv.Atoi(groups[2])
			if err == nil {
				return groups[1], itemNumber, true
			}
		}
	}

	return "", 0, false
}

func IsBelongsToSet(customID string) bool {
	return setCustomID.MatchString(customID) ||
		setBonusCustomID.MatchString(customID)
}
