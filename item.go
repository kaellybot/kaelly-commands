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

	itemRecipeCustomIDGroups = 2
)

var (
	itemCustomID       = regexp.MustCompile(fmt.Sprintf("^/%s$", ItemCommandName))
	itemRecipeCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/recipe$", ItemCommandName))
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

func CraftItemCustomID() string {
	return fmt.Sprintf("/%s", ItemCommandName)
}

func CraftItemRecipeCustomID(itemID string) string {
	return fmt.Sprintf("/%s/%s/recipe", ItemCommandName, itemID)
}

func ExtractItemRecipeCustomID(customID string) (string, bool) {
	if groups, ok := regex.ExtractCustomID(customID, itemRecipeCustomID,
		itemRecipeCustomIDGroups); ok {
		return groups[1], true
	}

	return "", false
}

func IsBelongsToItem(customID string) bool {
	return regex.IsBelongTo(customID, itemCustomID, itemRecipeCustomID)
}
