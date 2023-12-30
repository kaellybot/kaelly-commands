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
	HelpCommandName = "help"
)

var (
	helpCustomID = regexp.MustCompile(fmt.Sprintf("^/%s$", HelpCommandName))
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

func IsBelongsToHelp(customID string) bool {
	return regex.IsBelongTo(customID, helpCustomID)
}
