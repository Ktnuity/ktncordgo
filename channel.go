package ktncordgo

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (self *DiscordChannelUnit) Discord() IDiscordUnit {
	return self.discord
}

func (self *DiscordChannelUnit) Native() *discordgo.Channel {
	return self.channel
}

func (self *DiscordChannelUnit) Snowflake() string {
	return self.channel.ID
}

func (self *DiscordChannelUnit) Id() string {
	return self.channel.ID
}

func (self *DiscordChannelUnit) Guild() (IDiscordGuildUnit, error) {
	return self.discord.GetGuild(self.channel.GuildID)
}

func (self *DiscordChannelUnit) Name() string {
	return self.channel.Name
}

func (self *DiscordChannelUnit) Topic() string {
	return self.channel.Topic
}

func (self *DiscordChannelUnit) Position() int {
	return self.channel.Position
}

func (self *DiscordChannelUnit) NSFW() bool {
	return self.channel.NSFW
}

func (self *DiscordChannelUnit) Type() discordgo.ChannelType {
	return self.channel.Type
}

func (self *DiscordChannelUnit) Flags() discordgo.ChannelFlags {
	return self.channel.Flags
}

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

func (self *DiscordChannelUnit) FetchMessages(limit int) ([]IDiscordMessageUnit, error) {
	if limit > 100 {
		log.Printf("FetchMessages limit '%d' is larger than max allowed '%d'\n", limit, 100)
		limit = 100
	} else if limit < 1 {
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

func (self *DiscordChannelUnit) SendTyping() error {
	return self.discord.session.ChannelTyping(self.channel.ID)
}
