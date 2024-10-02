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
	HelpCommandName = "help"

	helpPageCustomIDGroups = 3
)

var (
	helpCustomID     = regexp.MustCompile(fmt.Sprintf("^/%s$", HelpCommandName))
	helpPageCustomID = regexp.MustCompile(fmt.Sprintf("^/%s/(\\w+)/details\\?page=(\\d+)$", HelpCommandName))
)

//nolint:nolintlint,exhaustive,lll,dupl
func getHelpSlashCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     HelpCommandName,
		Description:              "help.description",
		Type:                     discordgo.ChatApplicationCommand,
		DefaultMemberPermissions: constants.GetDefaultPermission(),
		DMPermission:             constants.GetDMPermission(),
		DescriptionLocalizations: i18n.GetLocalizations("help.description"),
	}
}

func CraftHelpCustomID() string {
	return fmt.Sprintf("/%s", HelpCommandName)
}

func CraftHelpPageCustomID(commandName string, page int) string {
	return fmt.Sprintf("/%s/%s/details?page=%v", HelpCommandName, commandName, page)
}

func ExtractHelpPageCustomID(customID string) (string, int, bool) {
	if groups, ok := regex.ExtractCustomID(customID, helpPageCustomID,
		helpPageCustomIDGroups); ok {
		commandName := groups[1]

		page, err := strconv.Atoi(groups[2])
		if err != nil {
			return "", -1, false
		}

		return commandName, page, true
	}

	return "", -1, false
}

func IsBelongsToHelp(customID string) bool {
	return regex.IsBelongTo(customID, helpCustomID, helpPageCustomID)
}
