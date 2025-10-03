package ktncordgo

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Discord returns the parent [DiscordUnit] object, the root of [ktncordgo].
//
// See: [DiscordUnit]
func (self *DiscordInteractionUnit) Discord() IDiscordUnit {
	return self.discord
}

// Native returns the underlying [discordgo.User] object.
//
// See: [discordgo.InteractionCreate]
func (self *DiscordInteractionUnit) Native() *discordgo.InteractionCreate {
	return self.interaction
}

// User returns the [DiscordUserUnit] instance of the command sender.
func (self *DiscordInteractionUnit) User() IDiscordUserUnit {
	return &DiscordUserUnit{
		discord: self.discord,
		user: self.interaction.User,
	}
}

// DeferReply defers the reply, giving more time before replying to the command.
//
// Typically discord requires a response within 3 seconds. Defer allows a delay of this.
// This cannot be followed by a [Reply] call and instead requires a [EditReply] call to follow up.
//
// Returns an error on failure.
//
// See: [discordgo.Session.InteractionRespond]
// See: [discordgo.InteractionCreate.Interaction]
// See: [discordgo.InteractionResponse]
// See: [discordgo.InteractionResponseDeferredChannelMessageWithSource]
func (self *DiscordInteractionUnit) DeferReply() error {
	return self.discord.session.InteractionRespond(self.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
}

// Reply sends a reply to user interaction. This cannot be used with [DeferReply].
//
// Parameters:
//   message - The text to include in the interaction response.
//
// Returns an error on failure.
//
// See: [discordgo.Session.InteractionRespond]
// See: [discordgo.InteractionCreate.Interaction]
// See: [discordgo.InteractionResponse]
// See: [discordgo.InteractionResponseData]
// See: [discordgo.InteractionResponseChannelMessageWithSource]
func (self *DiscordInteractionUnit) Reply(message string) error {
	return self.discord.session.InteractionRespond(self.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}

// EditReply edits the interaction reply, or sends it if [DeferReply] was used.
//
// Parameters:
//   message - Reference to the new content to use when editing response.
//
// Returns an error on failure.
//
// See: [discordgo.Session.InteractionResponseEdit]
// See: [discordgo.InteractionCreate.Interaction]
// See: [discordgo.WebhookEdit]
func (self *DiscordInteractionUnit) EditReply(message *string) error {
	_, err := self.discord.session.InteractionResponseEdit(self.interaction.Interaction, &discordgo.WebhookEdit{
		Content: message,
	})

	return err
}

// CommandName returns the name/label of the slash command.
//
// See: [discordgo.InteractionCreate.ApplicationCommandData]
// See: [discordgo.ApplicationCommandInteractionData.Name]
func (self *DiscordInteractionUnit) CommandName() string {
	return self.interaction.ApplicationCommandData().Name
}

// IsCommandName returns true if the name/label of the slash command matches a provided value.
//
// Parameters:
//   name - The name of the command to test against.
//
// Returns true if the command name matches.
//
// See: [discordgo.InteractionCreate.ApplicationCommandData]
// See: [discordgo.ApplicationCommandInteractionData.Name]
func (self *DiscordInteractionUnit) IsCommandName(name string) bool {
	return self.interaction.ApplicationCommandData().Name == name
}

// DispatchEvent matches the command name to a provided value, and if valid runs a provided callback.
//
// Parameters:
//   name - The name of the command to test against.
//   callback - The command handler for the given command.
//
// Returns true if the command matched.
//
// See: [IDiscordCommandFn]
// See: [IsCommandName]
func (self *DiscordInteractionUnit) DispatchEvent(name string, callback IDiscordCommandFn) bool {
	if !self.IsCommandName(name) {
		return false
	}

	err := callback(self)
	if err != nil {
		log.Printf("Failed to run dispatched event '%s': %v\n", name, err)
	}

	return true
}
