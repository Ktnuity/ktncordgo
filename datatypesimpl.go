package ktncordgo

import (
	"time"

	"github.com/ktnuity/ktnuitygo"
	"github.com/bwmarrin/discordgo"
)

// Build turns [DiscordMessageSend] into [discordgo.MessageSend].
func (self *DiscordMessageSend) Build() *discordgo.MessageSend {
	if self == nil { return nil }

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

func (self *DiscordMessageSend) ToEdit() *DiscordMessageEdit {
	if self == nil { return nil }

	return &DiscordMessageEdit{
		Content: &self.Content,
		Embeds: &self.Embeds,
		AllowedMentions: self.AllowedMentions,
	}
}

// Build turns [DiscordMessageEdit] into [discordgo.MessageEdit].
func (self *DiscordMessageEdit) Build() *discordgo.MessageEdit {
	if self == nil { return nil }
	var embeds *[]*discordgo.MessageEmbed = nil

	if self.Embeds != nil {
		newEmbeds := convertAll(*self.Embeds, func (embed *DiscordEmbed) *discordgo.MessageEmbed {
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

func (self *DiscordMessageEdit) ToSend() *DiscordMessageSend {
	if self == nil { return nil }

	content := ""
	if self.Content != nil { content = *self.Content }

	embeds := make([]*DiscordEmbed, 0, 4)
	if self.Embeds != nil {
		for _, e := range *self.Embeds {
			embeds = append(embeds, e)
		}
	}

	return &DiscordMessageSend{
		Content: content,
		Embeds: embeds,
		TTS: false,
		Attachments: make([]*DiscordAttachment, 0),
		AllowedMentions: self.AllowedMentions,
		Reference: nil,
	}
}

// Build turns [DiscordEmbed] into [discordgo.MessageEmbed].
func (self *DiscordEmbed) Build() *discordgo.MessageEmbed {
	if self == nil { return nil }
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

// Build turns [DiscordEmbedFooter] into [discordgo.MessageEmbedFooter].
func (self *DiscordEmbedFooter) Build() *discordgo.MessageEmbedFooter {
	if self == nil { return nil }
	return &discordgo.MessageEmbedFooter{
		Text: self.Text,
		IconURL: self.IconURL,
	}
}

// Build turns [DiscordEmbedImage] into [discordgo.MessageEmbedImage].
func (self *DiscordEmbedImage) Build() *discordgo.MessageEmbedImage {
	if self == nil { return nil }
	return &discordgo.MessageEmbedImage{
		URL: self.URL,
	}
}

// Build turns [DiscordEmbedField] into [discordgo.MessageEmbedField].
func (self *DiscordEmbedField) Build() *discordgo.MessageEmbedField {
	if self == nil { return nil }
	return &discordgo.MessageEmbedField{
		Name: self.Name,
		Value: self.Value,
		Inline: self.Inline,
	}
}

// Build turns [DiscordAttachment] into [discordgo.File].
func (self *DiscordAttachment) Build() *discordgo.File {
	if self == nil { return nil }
	return &discordgo.File{
		Name: self.Name,
		Reader: self.Source,
	}
}

// Build turns [DiscordAllowedMentions] into [discordgo.MessageAllowedMentions].
func (self *DiscordAllowedMentions) Build() *discordgo.MessageAllowedMentions {
	if self == nil { return nil }

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

// convertAll maps a slice of [T] into a slice of [U] using a mapper function.
func convertAll[T any, U any](list []T, fn func(T)U) []U {
	result := make([]U, len(list))

	for i, value := range list {
		result[i] = fn(value)
	}

	return result
}
