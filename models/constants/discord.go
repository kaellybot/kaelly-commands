package constants

import "github.com/bwmarrin/discordgo"

// GetContexts restricts commands to guilds (no bot DM, no private channel).
func GetContexts() *[]discordgo.InteractionContextType {
	return &[]discordgo.InteractionContextType{discordgo.InteractionContextGuild}
}

func GetDefaultPermission() *int64 {
	var defaultPermission int64 = discordgo.PermissionViewChannel
	return &defaultPermission
}

func GetManageServerPermission() *int64 {
	var manageServerPermission int64 = discordgo.PermissionManageGuild
	return &manageServerPermission
}
