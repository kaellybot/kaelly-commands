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
	ItemCommandName = "item"

	ItemQueryOptionName = "query"

	itemCustomIDGroups        = 2
	itemEffectsCustomIDGroups = 3
	itemRecipeCustomIDGroups  = 3
)

var (
	itemCustomID        = regexp.MustCompile(fmt.Sprintf("^/%s\\?type=(\\w+)$", ItemCommandName))
	itemEffectsCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/effects\\?type=(\\w+)$", ItemCommandName))
	itemRecipeCustomID  = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/recipe\\?type=(\\w+)$", ItemCommandName))
)

//nolint:exhaustive,lll,funlen
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

func CraftItemCustomID(itemType string) string {
	return fmt.Sprintf("/%s?type=%s", ItemCommandName, itemType)
}

func CraftItemEffectsCustomID(itemID, itemType string) string {
	return fmt.Sprintf("/%s/%s/effects?type=%s", ItemCommandName, itemID, itemType)
}

func CraftItemRecipeCustomID(itemID, itemType string) string {
	return fmt.Sprintf("/%s/%s/recipe?type=%s", ItemCommandName, itemID, itemType)
}

func ExtractItemCustomID(customID string) (string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, itemCustomID,
		itemCustomIDGroups); ok {
		return groups[1], true
	}

	return "", false
}

func ExtractItemEffectsCustomID(customID string) (string, string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, itemEffectsCustomID,
		itemEffectsCustomIDGroups); ok {
		return groups[1], groups[2], true
	}

	return "", "", false
}

func ExtractItemRecipeCustomID(customID string) (string, string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, itemRecipeCustomID,
		itemRecipeCustomIDGroups); ok {
		return groups[1], groups[2], true
	}

	return "", "", false
}

func IsBelongsToItem(customID string) bool {
	return regex.IsBelongTo(customID, itemCustomID,
		itemEffectsCustomID, itemRecipeCustomID)
}
