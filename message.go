package ktncordgo

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ktnuity/ktnuitygo"
)

func (self *DiscordMessageUnit) Discord() IDiscordUnit {
	return self.discord
}

func (self *DiscordMessageUnit) Native() *discordgo.Message {
	return self.message
}

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

func (self *DiscordMessageUnit) Author() IDiscordUserUnit {
	if self.message.Author == nil {
		return nil
	}

	return &DiscordUserUnit{
		discord: self.discord,
		user: self.message.Author,
	}
}

func (self *DiscordMessageUnit) Timestamp() time.Time {
	return self.message.Timestamp
}

func (self *DiscordMessageUnit) EditedTimestamp() *time.Time {
	return self.message.EditedTimestamp
}

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

func (self *DiscordMessageUnit) Edit(message string) error {
	msg, err := self.discord.session.ChannelMessageEdit(self.message.ChannelID, self.message.ID, message)

	if err != nil {
		return err
	}

	self.message = msg
	return nil
}

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

func (self *DiscordMessageUnit) Crosspost() error {
	msg, err := self.discord.session.ChannelMessageCrosspost(self.message.ChannelID, self.message.ID)

	if err != nil {
		return err
	}

	self.message = msg
	return nil
}

func (self *DiscordMessageUnit) Delete() error {
	return self.discord.session.ChannelMessageDelete(self.message.ChannelID, self.message.ID)
}

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
