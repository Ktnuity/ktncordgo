package ktncordgo

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Discord returns the parent [DiscordUnit] object, the root of [ktncordgo].
//
// See: [DiscordUnit]
func (self *DiscordChannelUnit) Discord() IDiscordUnit {
	return self.discord
}

// Native returns the underlying [discordgo.Channel] object.
//
// See: [discordgo.Channel]
func (self *DiscordChannelUnit) Native() *discordgo.Channel {
	return self.channel
}

// Snowflake returns the ID of the discord channel.
//
// See: [discordgo.Channel.ID]
func (self *DiscordChannelUnit) Snowflake() string {
	return self.channel.ID
}

// Id returns the ID of the discord channel.
//
// See: [DiscordChannelUnit.Snowflake]
func (self *DiscordChannelUnit) Id() string {
	return self.channel.ID
}

// Guild returns the Guild that the current channel is in.
//
// See: [DiscordGuildUnit]
// See: [DiscordUnit.GetGuild]
func (self *DiscordChannelUnit) Guild() (IDiscordGuildUnit, error) {
	return self.discord.GetGuild(self.channel.GuildID)
}

// Name returns the name of the current channel.
//
// See: [discordgo.Channel.Name]
func (self *DiscordChannelUnit) Name() string {
	return self.channel.Name
}

// Topic returns the topic of the current channel.
//
// See: [discordgo.Channel.Topic]
func (self *DiscordChannelUnit) Topic() string {
	return self.channel.Topic
}

// Position returns the position of the current channel in the channel list.
//
// See: [discordgo.Channel.Position]
func (self *DiscordChannelUnit) Position() int {
	return self.channel.Position
}

// NSFW returns true if the channel is marked as a NSFW channel.
//
// See: [discordgo.Channel.NSFW]
func (self *DiscordChannelUnit) NSFW() bool {
	return self.channel.NSFW
}

// Type returns the type of the channel.
//
// See: [discordgo.ChannelType]
// See: [discordgo.Channel.Type]
func (self *DiscordChannelUnit) Type() discordgo.ChannelType {
	return self.channel.Type
}

// Flags returns the flags of the channel.
//
// See: [discordgo.ChannelFlags]
// See: [discordgo.Channel.Flags]
func (self *DiscordChannelUnit) Flags() discordgo.ChannelFlags {
	return self.channel.Flags
}

// FetchMessage finds and returns a single message sent in the channel.
//
// Parameters:
//   messageId - The ID of the message to look for.
//
// Returns the [DiscordMessageUnit] of the message if found, otherwise an error.
//
// See: [DiscordMessageUnit]
// See: [discordgo.Session.ChannelMessage]
func (self *DiscordChannelUnit) FetchMessage(messageId string) (IDiscordMessageUnit, error) {
	msg, err := self.discord.session.ChannelMessage(self.channel.ID, messageId)
	if err != nil {
		return nil, err
	}

	return &DiscordMessageUnit{
		discord: self.discord,
		message: msg,
	}, nil
}

// FetchMessages finds and returns the latest X messages in the channel.
//
// Parameters:
//   limit - The amount of channels to find. Min: 1. Max: 100.
//
// Returns a slice of the found messages on success, otherwise an error.
//
// See: [DiscordMessageUnit]
// See: [discordgo.Session.ChannelMessages]
func (self *DiscordChannelUnit) FetchMessages(limit int) ([]IDiscordMessageUnit, error) {
	if limit > 100 {
		log.Printf("FetchMessages limit '%d' is larger than max allowed '%d'\n", limit, 100)
		limit = 100
	} else if limit < 1 { // I know minimul is mentioned to be 1, but we're not just gonna error if the user pick anything less.
		log.Printf("FetchMesssages limit '%d' is less than min allowed '%d'\n", limit, 0)
		limit = 0
	}

	if limit == 0 {
		return []IDiscordMessageUnit{}, nil
	}

	messages, err := self.discord.session.ChannelMessages(self.channel.ID, limit, "", "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch channel messages: %v", err)
	}

	result := make([]IDiscordMessageUnit, len(messages))

	for i, message := range messages {
		result[i] = &DiscordMessageUnit{
			discord: self.discord,
			message: message,
		}
	}

	return result, nil
}

// GetLastMassage finds and returns the latest message in the channel.
//
// Returns the message if found, otherwise an error.
//
// See: [DiscordMessageUnit]
// See: [DiscordChannelUnit.FetchMessages]
func (self *DiscordChannelUnit) GetLastMessage() (IDiscordMessageUnit, error) {
	messages, err := self.FetchMessages(100)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch last message: %v", err)
	}

	if len(messages) == 0 {
		return nil, fmt.Errorf("failed to fetch last message: no messages found")
	}

	return messages[0], nil
}

// SendMessage sends a message in the channel.
//
// Parameters:
//   message - The content of the message to send.
//
// Returns the sent message on success, otherwise an error.
//
// See: [DiscordMessageUnit]
// See: [discordgo.Session.ChannelMessageSend]
func (self *DiscordChannelUnit) SendMessage(message string) (IDiscordMessageUnit, error) {
	msg, err := self.discord.session.ChannelMessageSend(self.channel.ID, message)
	if err != nil {
		return nil, err
	}

	return &DiscordMessageUnit{
		discord: self.discord,
		message: msg,
	}, nil
}

// SendMessageOptions sends a message in the channel with options.
//
// Parameters:
//   options - The message options for the message to send.
//
// Returns the sent message on success, otherwise an error.
//
// See: [DiscordMessageUnit]
// See: [discordgo.Session.ChannelMessageSendComplex]
func (self *DiscordChannelUnit) SendMessageOptions(options DiscordMessageSend) (IDiscordMessageUnit, error) {
	msg, err := self.discord.session.ChannelMessageSendComplex(self.channel.ID, options.Build())
	if err != nil {
		return nil, fmt.Errorf("failed to send message with options: %v", err)
	}

	return &DiscordMessageUnit{
		discord: self.discord,
		message: msg,
	}, nil
}

// SendTyping sends the typing animation to users in the channel.
//
// Returns an error on failure.
//
// See: [discordgo.Session.ChannelTyping]
func (self *DiscordChannelUnit) SendTyping() error {
	return self.discord.session.ChannelTyping(self.channel.ID)
}
