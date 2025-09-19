package ktncordgo

import "github.com/bwmarrin/discordgo"

func (self *DiscordUserUnit) Discord() IDiscordUnit {
	return self.discord
}

func (self *DiscordUserUnit) Native() *discordgo.User {
	return self.user
}

func (self *DiscordUserUnit) Snowflake() string {
	return self.user.ID
}

func (self *DiscordUserUnit) Id() string {
	return self.user.ID
}

func (self *DiscordUserUnit) Username() string {
	return self.user.Username
}

func (self *DiscordUserUnit) Discriminator() string {
	return self.user.Discriminator
}

func (self *DiscordUserUnit) GlobalName() string {
	return self.user.GlobalName
}

func (self *DiscordUserUnit) IsBot() bool {
	return self.user.Bot
}

func (self *DiscordUserUnit) IsVerified() bool {
	return self.user.Verified
}

func (self *DiscordUserUnit) HasMFAEnabled() bool {
	return self.user.MFAEnabled
}

func (self *DiscordUserUnit) IsSystem() bool {
	return self.user.System
}

func (self *DiscordUserUnit) IsAnyNitro() bool {
	return self.user.PremiumType != discordgo.UserPremiumTypeNone
}

func (self *DiscordUserUnit) IsNitro() bool {
	return self.user.PremiumType == discordgo.UserPremiumTypeNitro
}

func (self *DiscordUserUnit) IsNitroClassic() bool {
	return self.user.PremiumType == discordgo.UserPremiumTypeNitroClassic
}

func (self *DiscordUserUnit) IsNitroBasic() bool {
	return self.user.PremiumType == discordgo.UserPremiumTypeNitroBasic
}
