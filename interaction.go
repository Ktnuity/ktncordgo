package ktncordgo

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (self *DiscordInteractionUnit) Discord() IDiscordUnit {
	return self.discord
}

func (self *DiscordInteractionUnit) Native() *discordgo.InteractionCreate {
	return self.interaction
}

func (self *DiscordInteractionUnit) DeferReply() error {
	return self.discord.session.InteractionRespond(self.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
}

func (self *DiscordInteractionUnit) Reply(message string) error {
	return self.discord.session.InteractionRespond(self.interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}

func (self *DiscordInteractionUnit) EditReply(message *string) error {
	_, err := self.discord.session.InteractionResponseEdit(self.interaction.Interaction, &discordgo.WebhookEdit{
		Content: message,
	})

	return err
}

func (self *DiscordInteractionUnit) CommandName() string {
	return self.interaction.ApplicationCommandData().Name
}

func (self *DiscordInteractionUnit) IsCommandName(name string) bool {
	return self.interaction.ApplicationCommandData().Name == name
}

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
