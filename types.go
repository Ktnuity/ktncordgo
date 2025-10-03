package ktncordgo

import "github.com/bwmarrin/discordgo"

// DiscordUnit holds the main instance of ktncordgo.
//
// See: [discordgo.Session]
type DiscordUnit struct {
	session *discordgo.Session
}

// DiscordInteractionUnit holds any interaction related functionality,
// typically seen with interaction events, like slash commands.
//
// See: [discordgo.InteractionCreate]
type DiscordInteractionUnit struct {
	discord *DiscordUnit
	interaction *discordgo.InteractionCreate
}

// DiscordGuildUnit is the wrapper for the Guild object
//
// See: [discordgo.Guild]
type DiscordGuildUnit struct {
	discord *DiscordUnit
	guild *discordgo.Guild
}

// DiscordChannelUnit is the wrapper for the Channel object
//
// See: [discordgo.Channel]
type DiscordChannelUnit struct {
	discord *DiscordUnit
	channel *discordgo.Channel
}

// DiscordMessageUnit is the wrapper for the Message object
//
// See: [discordgo.Message]
type DiscordMessageUnit struct {
	discord *DiscordUnit
	message *discordgo.Message
}

// DiscordUserUnit is the wrapper for the User object
//
// See: [discordgo.User]
type DiscordUserUnit struct {
	discord *DiscordUnit
	user *discordgo.User
}
