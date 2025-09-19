package ktncordgo

import (
	"time"

	"github.com/ktnuity/ktnuitygo"
	"github.com/bwmarrin/discordgo"
)

func (self *DiscordMessageSend) Build() *discordgo.MessageSend {
	var reference *discordgo.MessageReference = nil
	if self.Reference != nil {
		reference = &discordgo.MessageReference{
			Type: discordgo.MessageReferenceTypeDefault,
			MessageID: self.Reference.Native().ID,
			ChannelID: self.Reference.Native().ChannelID,
			GuildID: self.Reference.Native().GuildID,
			FailIfNotExists: ktnuitygo.AsRef(true),
		}
	}

	return &discordgo.MessageSend{
		Content: self.Content,
		Embeds: convertAll(self.Embeds, func (embed *DiscordEmbed) *discordgo.MessageEmbed {
			return embed.Build()
		}),
		TTS: self.TTS,
		Files: convertAll(self.Attachments, func (attachment *DiscordAttachment) *discordgo.File {
			return attachment.Build()
		}),
		AllowedMentions: self.AllowedMentions.Build(),
		Reference: reference,
	}
}

func (self *DiscordMessageEdit) Build() *discordgo.MessageEdit {
	var embeds *[]*discordgo.MessageEmbed = nil

	if self.embeds != nil {
		newEmbeds := convertAll(*self.embeds, func (embed *DiscordEmbed) *discordgo.MessageEmbed {
			return embed.Build()
		})
		embeds = &newEmbeds
	}

	var mentions *discordgo.MessageAllowedMentions = nil
	if self.AllowedMentions != nil {
		self.AllowedMentions.Build()
	}

	return &discordgo.MessageEdit{
		Content: self.Content,
		Embeds: embeds,
		AllowedMentions: mentions,
	}
}

func (self *DiscordEmbed) Build() *discordgo.MessageEmbed {
	timestamp := ""
	if self.Timestamp != nil {
		timestamp = self.Timestamp.Format(time.RFC3339)
	}

	return &discordgo.MessageEmbed{
		URL: self.URL,
		Title: self.Title,
		Description: self.Description,
		Timestamp: timestamp,
		Color: self.Color,
		Footer: self.Footer.Build(),
		Image: self.Image.Build(),
		Fields: convertAll(self.Fields, func (field *DiscordEmbedField) *discordgo.MessageEmbedField {
			return field.Build()
		}),
	}
}

func (self *DiscordEmbedFooter) Build() *discordgo.MessageEmbedFooter {
	return &discordgo.MessageEmbedFooter{
		Text: self.Text,
		IconURL: self.IconURL,
	}
}

func (self *DiscordEmbedImage) Build() *discordgo.MessageEmbedImage {
	return &discordgo.MessageEmbedImage{
		URL: self.URL,
	}
}

func (self *DiscordEmbedField) Build() *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name: self.Name,
		Value: self.Value,
		Inline: self.Inline,
	}
}

func (self *DiscordAttachment) Build() *discordgo.File {
	return &discordgo.File{
		Name: self.Name,
		Reader: self.Source,
	}
}

func (self *DiscordAllowedMentions) Build() *discordgo.MessageAllowedMentions {
	result := &discordgo.MessageAllowedMentions{}
	parse := make([]discordgo.AllowedMentionType, 0, 3)

	if self.Users != nil {
		result.Users = self.Users
		parse = append(parse, discordgo.AllowedMentionTypeUsers)
	}

	if self.Roles != nil {
		result.Roles = self.Roles
		parse = append(parse, discordgo.AllowedMentionTypeRoles)
	}

	if len(parse) == 0 {
		parse = append(parse, discordgo.AllowedMentionTypeEveryone)
	}

	result.Parse = parse

	return result
}

func convertAll[T any, U any](list []T, fn func(T)U) []U {
	result := make([]U, len(list))

	for i, value := range list {
		result[i] = fn(value)
	}

	return result
}
