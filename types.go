package ktncordgo

import "github.com/bwmarrin/discordgo"

type DiscordUnit struct {
	session *discordgo.Session
}

type DiscordInteractionUnit struct {
	discord *DiscordUnit
	interaction *discordgo.InteractionCreate
}

type DiscordGuildUnit struct {
	discord *DiscordUnit
	guild *discordgo.Guild
}

type DiscordChannelUnit struct {
	discord *DiscordUnit
	channel *discordgo.Channel
}

type DiscordMessageUnit struct {
	discord *DiscordUnit
	message *discordgo.Message
}

type DiscordUserUnit struct {
	discord *DiscordUnit
	user *discordgo.User
}
