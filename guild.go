package ktncordgo

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Discord returns the parent [DiscordUnit] object, the root of [ktncordgo].
//
// See: [DiscordUnit]
func (self *DiscordGuildUnit) Discord() IDiscordUnit {
	return self.discord
}

// Native returns the underlying [discordgo.Guild] object.
//
// See: [discordgo.Guild]
func (self *DiscordGuildUnit) Native() *discordgo.Guild {
	return self.guild
}

// Snowflake returns the ID of the discord guild.
//
// See: [discordgo.Guild.ID]
func (self *DiscordGuildUnit) Snowflake() string {
	return self.guild.ID
}

// Id returns the ID of the discord guild.
//
// See: [DiscordGuildUser.Snowflake]
func (self *DiscordGuildUnit) Id() string {
	return self.guild.ID
}

// Name returns the name of the discord guild.
//
// See: [discordgo.Guild.Name]
func (self *DiscordGuildUnit) Name() string {
	return self.guild.Name
}

// Description returns the title/description of the discord guild.
//
// See: [discordgo.Guild.Description]
func (self *DiscordGuildUnit) Description() string {
	return self.guild.Description
}

// Icon returns the icon URL of the discord guild.
//
// See: [discordgo.Guild.Icon]
func (self *DiscordGuildUnit) Icon() string {
	return self.guild.Icon
}

// Region returns the voice channel region used in the discord guild.
//
// See: [discordgo.Guild.Region]
func (self *DiscordGuildUnit) Region() string {
	return self.guild.Region
}

// IsOwner returns [true] if the bot user of the underlying [DiscordUnit] is the owner of the discord guild.
//
// See: [discordgo.Guild.Owner]
func (self *DiscordGuildUnit) IsOwner() bool {
	return self.guild.Owner
}

// GetChannels returns the channels available in the discord guild.
// Returns an error if there's an issure with finding channels.
//
// See [DiscordChannelUnit]
// See: [discordgo.Session.GuildChannels]
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

// GetChannel returns a channel in the discord guild.
//
// Parameters
//   channelId - The ID of the discord channel.
//
// Returns the channel if found, otherwise an error.
//
// See: [DiscordChannelUnit]
// See: [DiscordGuildUnit.GetChannels]
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

// GetMember Count returns the number of members in the discord guild.
// Note: This is an iterative process, be vary of usage in larger servers.
//
// Returns number of members if successful, otherwise an error.
//
// See: [discordgo.Session.GuildMembers]
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
