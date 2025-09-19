package ktncordgo

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (self *DiscordGuildUnit) Discord() IDiscordUnit {
	return self.discord
}

func (self *DiscordGuildUnit) Native() *discordgo.Guild {
	return self.guild
}

func (self *DiscordGuildUnit) Snowflake() string {
	return self.guild.ID
}

func (self *DiscordGuildUnit) Id() string {
	return self.guild.ID
}

func (self *DiscordGuildUnit) Name() string {
	return self.guild.Name
}

func (self *DiscordGuildUnit) Description() string {
	return self.guild.Description
}

func (self *DiscordGuildUnit) Icon() string {
	return self.guild.Icon
}

func (self *DiscordGuildUnit) Region() string {
	return self.guild.Region
}

func (self *DiscordGuildUnit) IsOwner() bool {
	return self.guild.Owner
}

func (self *DiscordGuildUnit) GetChannels() ([]IDiscordChannelUnit, error) {
	chans, err := self.discord.session.GuildChannels(self.guild.ID)
	if err != nil {
		return nil, err
	}

	result := make([]IDiscordChannelUnit, len(chans))

	for i, channel := range chans {
		result[i] = &DiscordChannelUnit{
			discord: self.discord,
			channel: channel,
		}
	}

	return result, nil
}

func (self *DiscordGuildUnit) GetChannel(channelId string) (IDiscordChannelUnit, error) {
	channels, err := self.GetChannels()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch channel: %v", err)
	}

	for _, channel := range channels {
		if channel.Id() == channelId {
			return channel, nil
		}
	}

	return nil, fmt.Errorf("failed to fetch channel: channel not found")
}

func (self *DiscordGuildUnit) GetMemberCount() (int, error) {
	var count int = 0
	var last string = ""

	log.Printf("Scanning member count for '%s'[%s]\n", self.guild.Name, self.guild.ID)

	for {
		membs, err := self.discord.session.GuildMembers(self.guild.ID, last, 1000)
		if err != nil {
			return 0, err
		}

		next := len(membs)
		count += next

		log.Printf("New count for [%s]: %d\n", self.guild.ID, count)

		if next < 1000 {
			break
		}

		last = membs[len(membs) - 1].User.ID
	}

	log.Printf("Got member count for [%s]: %d\n", self.guild.ID, count)

	return count, nil
}
