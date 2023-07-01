package commands

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/kaellybot/kaelly-commands/models/constants"
	i18n "github.com/kaysoro/discordgo-i18n"
)

const (
	ItemCommandName = "item"

	ItemQueryOptionName = "query"

	itemCustomIDGroups       = 2
	itemRecipeCustomIDGroups = 2
)

var (
	itemCustomID       = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)$", ItemCommandName))
	itemRecipeCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/recipe$", ItemCommandName))
)

//nolint:nolintlint,exhaustive,lll,dupl,funlen
func getItemSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     ItemCommandName,
		Description:              "item.description",
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("item.description"),
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:                     ItemQueryOptionName,
				Description:              "item.query.description",
				NameLocalizations:        *i18n.GetLocalizations("item.query.name"),
				DescriptionLocalizations: *i18n.GetLocalizations("item.query.description"),
				Type:                     discordgo.ApplicationCommandOptionString,
				Required:                 true,
				Autocomplete:             true,
			},
		},
	}
}

func CraftItemCustomID(itemID string) string {
	return fmt.Sprintf("/%s/%s", ItemCommandName, itemID)
}

func CraftItemRecipeCustomID(itemID string) string {
	return fmt.Sprintf("/%s/%s/recipe", ItemCommandName, itemID)
}

func ExtractItemCustomID(customID string) (string, bool) {
	if itemCustomID.MatchString(customID) {
		groups := itemCustomID.FindStringSubmatch(customID)
		if len(groups) == itemCustomIDGroups {
			return groups[1], true
		}
	}

	return "", false
}

func ExtractItemRecipeCustomID(customID string) (string, bool) {
	if itemRecipeCustomID.MatchString(customID) {
		groups := itemRecipeCustomID.FindStringSubmatch(customID)
		if len(groups) == itemRecipeCustomIDGroups {
			return groups[1], true
		}
	}

	return "", false
}

func IsBelongsToItem(customID string) bool {
	return itemCustomID.MatchString(customID) ||
		itemRecipeCustomID.MatchString(customID)
}
