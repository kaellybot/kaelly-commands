package commands

import "github.com/bwmarrin/discordgo"

func GetDMPermission() *bool {
	var dmPermission = false
	return &dmPermission
}

func GetDefaultPermission() *int64 {
	var defaultPermission int64 = discordgo.PermissionViewChannel
	return &defaultPermission
}

func GetManageServerPermission() *int64 {
	var manageServerPermission int64 = discordgo.PermissionManageServer
	return &manageServerPermission
}
