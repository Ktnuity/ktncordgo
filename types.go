package ktncordgo

import "github.com/bwmarrin/discordgo"

// DiscordUnit holds the main instance of ktncordgo.
type DiscordUnit struct {
	session *discordgo.Session
}

// DiscordInteractionUnit holds any interaction related functionality,
// typically seen with interaction events, like slash commands.
type DiscordInteractionUnit struct {
	discord *DiscordUnit
	interaction *discordgo.InteractionCreate
}

// DiscordGuildUnit is the wrapper for the Guild object
type DiscordGuildUnit struct {
	discord *DiscordUnit
	guild *discordgo.Guild
}

// DiscordChannelUnit is the wrapper for the Channel object
type DiscordChannelUnit struct {
	discord *DiscordUnit
	channel *discordgo.Channel
}

// DiscordMessageUnit is the wrapper for the Message object
type DiscordMessageUnit struct {
	discord *DiscordUnit
	message *discordgo.Message
}

// DiscordUserUnit is the wrapper for the User object
type DiscordUserUnit struct {
	discord *DiscordUnit
	user *discordgo.User
}
