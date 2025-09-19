package ktncordgo

import (
	"io"
	"time"
)

type DiscordMessageSend struct {
	Content string
	Embeds []*DiscordEmbed
	TTS bool
	Attachments []*DiscordAttachment
	AllowedMentions *DiscordAllowedMentions
	Reference IDiscordMessageUnit
}

type DiscordMessageEdit struct {
	Content *string
	embeds *[]*DiscordEmbed
	AllowedMentions *DiscordAllowedMentions
}

type DiscordEmbed struct {
	URL string
	Title string
	Description string
	Timestamp *time.Time
	Color int
	Footer *DiscordEmbedFooter
	Image *DiscordEmbedImage
	Fields []*DiscordEmbedField
}

type DiscordEmbedFooter struct {
	Text string
	IconURL string
}

type DiscordEmbedImage struct {
	URL string
}

type DiscordEmbedField struct {
	Name string
	Value string
	Inline bool
}

type DiscordAttachment struct {
	Name string
	Source io.Reader
}

type DiscordMentionType string

const (
	DiscordMentionTypeRoles		DiscordMentionType = "roles"
	DiscordMentionTypeUsers		DiscordMentionType = "users"
	DiscordMentionTypeEveryone	DiscordMentionType = "everyone"
)

type DiscordAllowedMentions struct {
	Roles []string
	Users []string
	Replieduser bool
}

