package ktncordgo

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ktnuity/ktnuitygo"
)

// Discord returns the parent [DiscordUnit] object, the root of [ktncordgo].
func (self *DiscordMessageUnit) Discord() IDiscordUnit {
	return self.discord
}

// Native returns the underlying [discordgo.Message] object.
func (self *DiscordMessageUnit) Native() *discordgo.Message {
	return self.message
}

// Channel returns the channel the message was sent in.
//
// See: [DiscordChannelUnit]
func (self *DiscordMessageUnit) Channel() IDiscordChannelUnit {
	channel, err := self.discord.session.Channel(self.message.ChannelID)
	if err != nil {
		log.Printf("Failed to fetch channel (id '%s') from message (id '%s'): %v\n", self.message.ChannelID, self.message.ID, err)
		return nil
	}

	return &DiscordChannelUnit{
		discord: self.discord,
		channel: channel,
	}
}

// Author returns the sender of the message.
//
// See: [DiscordUserUnit]
func (self *DiscordMessageUnit) Author() IDiscordUserUnit {
	if self.message.Author == nil {
		return nil
	}

	return &DiscordUserUnit{
		discord: self.discord,
		user: self.message.Author,
	}
}

// Timestamp returns the [time.Time] when the message was sent.
func (self *DiscordMessageUnit) Timestamp() time.Time {
	return self.message.Timestamp
}

// EditedTimestamp returns the [time.Time] when the message was sent.
// It returns nil if the message was never edited.
func (self *DiscordMessageUnit) EditedTimestamp() *time.Time {
	return self.message.EditedTimestamp
}

// Mentions returns the users that have been mentioned in the message.
//
// See: [DiscordUserUnit]
func (self *DiscordMessageUnit) Mentions() []IDiscordUserUnit {
	result := make([]IDiscordUserUnit, len(self.message.Mentions))

	for i, user := range self.message.Mentions {
		result[i] = &DiscordUserUnit{
			discord: self.discord,
			user: user,
		}
	}

	return result
}

// Edit edits the message with new text. Only works if the [DiscordUnit] object user represent the sender of this message.
//
// Parameter:
//   message - the message content to apply with the edit.
//
// Returns an error upon failure.
func (self *DiscordMessageUnit) Edit(message string) error {
	msg, err := self.discord.session.ChannelMessageEdit(self.message.ChannelID, self.message.ID, message)

	if err != nil {
		return err
	}

	self.message = msg
	return nil
}

// EditOptions edits the message with new content. Only works if the [DiscordUnit] object user represent the sender of this message.
//
// Parameter:
//   options - the message content to apply with the edit.
//
// Returns an error upon failure.
func (self *DiscordMessageUnit) EditOptions(options DiscordMessageEdit) error {
	opts := options.Build()

	opts.ID = self.message.ID
	opts.Channel = self.message.ChannelID

	msg, err := self.discord.session.ChannelMessageEditComplex(opts)
	if err != nil {
		return fmt.Errorf("failed to edit message with options: %v", err)
	}

	self.message = msg
	return nil
}

// Crosspost performs a crosspost for the message. This only works in announcement type discord channels.
//
// Returns an error upon failure.
func (self *DiscordMessageUnit) Crosspost() error {
	msg, err := self.discord.session.ChannelMessageCrosspost(self.message.ChannelID, self.message.ID)

	if err != nil {
		return err
	}

	self.message = msg
	return nil
}

// Delete deletes the current message. This requires permission to manage messages or that the underlying [DiscordUnit] object user sent this message.
//
// Returns an error upon failure.
func (self *DiscordMessageUnit) Delete() error {
	return self.discord.session.ChannelMessageDelete(self.message.ChannelID, self.message.ID)
}

// Reply replies to the message.
//
// Parameters:
//   message - the message content to use in the reply.
//
// Returns an error upon failure.
func (self *DiscordMessageUnit) Reply(message string) (IDiscordMessageUnit, error) {
	msg, err := self.discord.session.ChannelMessageSendReply(self.message.ChannelID, message, &discordgo.MessageReference{
		Type: discordgo.MessageReferenceTypeDefault,
		MessageID: self.message.ID,
		ChannelID: self.message.ChannelID,
		GuildID: self.message.GuildID,
		FailIfNotExists: ktnuitygo.AsRef(true),
	})

	if err != nil {
		return nil, err
	}

	return &DiscordMessageUnit{
		discord: self.discord,
		message: msg,
	}, nil
}
