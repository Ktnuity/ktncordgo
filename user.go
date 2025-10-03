package ktncordgo

import "github.com/bwmarrin/discordgo"

// Discord returns the parent [DiscordUnit] object, the root of [ktncordgo].
func (self *DiscordUserUnit) Discord() IDiscordUnit {
	return self.discord
}

// Native returns the underlying [discordgo.User] object.
func (self *DiscordUserUnit) Native() *discordgo.User {
	return self.user
}

// Snowflake returns the ID of the discord user.
func (self *DiscordUserUnit) Snowflake() string {
	return self.user.ID
}

// Id returns the ID of the discord user.
//
// See: [DiscordUserUnit.Snowflake]
func (self *DiscordUserUnit) Id() string {
	return self.user.ID
}

// Username returns the username of the discord user.
func (self *DiscordUserUnit) Username() string {
	return self.user.Username
}

// Discriminator returns the 4 Discriminator digs of the discord user.
// This is no longer used for normal users and is limited to bots.
// For a normal user this will return "0000".
func (self *DiscordUserUnit) Discriminator() string {
	return self.user.Discriminator
}

// GlobalName returns the global display name of the discord user.
func (self *DiscordUserUnit) GlobalName() string {
	return self.user.GlobalName
}

// IsBot returns true if the discord user is a bot.
func (self *DiscordUserUnit) IsBot() bool {
	return self.user.Bot
}

// IsVerified returns true if the discord user has verified their email.
// Might not work. See [discordgo.User.Verified] for more info.
func (self *DiscordUserUnit) IsVerified() bool {
	return self.user.Verified
}

// HasMFAEnabled returns true if the discord user has multi-factor authentication enabled.
// Might not work. See [discordgo.User.MFAEnabled] for more info.
func (self *DiscordUserUnit) HasMFAEnabled() bool {
	return self.user.MFAEnabled
}

// IsSystem returns true if the discord user is a system user, typically seen with Discord's own announcement messages.
func (self *DiscordUserUnit) IsSystem() bool {
	return self.user.System
}

// IsAnyNitro returns true if the discord user has any of the Nitro tier, be that Nitro (9.99€), Nitro Basic (2.99€) or Nitro Classic (4.99€ mostly obsolete).
func (self *DiscordUserUnit) IsAnyNitro() bool {
	return self.user.PremiumType != discordgo.UserPremiumTypeNone
}

// IsNitro returns true if the discord user has Nitro (9.99€).
func (self *DiscordUserUnit) IsNitro() bool {
	return self.user.PremiumType == discordgo.UserPremiumTypeNitro
}

// IsNitroClassic returns true if the discord user has Nitro Classic (4.99€ mostly obsolete).
func (self *DiscordUserUnit) IsNitroClassic() bool {
	return self.user.PremiumType == discordgo.UserPremiumTypeNitroClassic
}

// IsNitroBasic returns true if the discord user has Nitro Basic (2.99€).
func (self *DiscordUserUnit) IsNitroBasic() bool {
	return self.user.PremiumType == discordgo.UserPremiumTypeNitroBasic
}
