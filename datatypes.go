package ktncordgo

import (
	"io"
	"time"
)

// DiscordMessageSend contains options used for [DiscordChannelUnit.SendMessageOptions].
//
// See: [DiscordEmbed]
// See: [DiscordAttachment]
// See: [DiscordAllowedMentions]
// See: [DiscordMessageUnit]
type DiscordMessageSend struct {
	Content string
	Embeds []*DiscordEmbed
	TTS bool
	Attachments []*DiscordAttachment
	AllowedMentions *DiscordAllowedMentions
	Reference IDiscordMessageUnit
}

// DiscordMessageEdit contains options used for [DiscordMessageUnit.EditOptions].
//
// See: [DiscordEmbed]
// See: [DiscordAllowedMentions]
type DiscordMessageEdit struct {
	Content *string
	Embeds *[]*DiscordEmbed
	AllowedMentions *DiscordAllowedMentions
}

// DiscordEmbed contains options used for [DiscordMessageSend] and [DiscordMessageEdit].
//
// See: parent [DiscordMessageSend] or [DiscordMessageEdit]
// See: [time.Time]
// See: [DiscordEmbedFooter]
// See: [DiscordEmbedImage]
// See: [DiscordEmbedField]
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

// DiscordEmbedFooter contains options used for [DiscordEmbed].
//
// See: parent [DiscordEmbed]
type DiscordEmbedFooter struct {
	Text string
	IconURL string
}

// DiscordEmbedImage contains options used for [DiscordEmbed].
//
// See: parent [DiscordEmbed]
type DiscordEmbedImage struct {
	URL string
}

// DiscordEmbedField contains options used for [DiscordEmbed].
//
// See: parent [DiscordEmbed]
type DiscordEmbedField struct {
	Name string
	Value string
	Inline bool
}

// DiscordAttachment contains options used for [DiscordMessageSend].
//
// See: parent [DiscordMessageSend]
// See: [io.Reader]
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

// DiscordAllowedMentions contains options used for [DiscordMessageSend] and [DiscordMesageEdit].
//
// See: parent [DiscordMessageSend] or [DiscordMessageEdit]
type DiscordAllowedMentions struct {
	Roles []string
	Users []string
	Replieduser bool
}

