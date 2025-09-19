package ktncordgo

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

type IDiscordUnit interface {
	Session() *discordgo.Session
	NewInteractionUnit(interaction *discordgo.InteractionCreate) IDiscordInteractionUnit

	OnSlashCommand(func (IDiscordUnit, IDiscordInteractionUnit))
	OnMessageCreate(func (IDiscordUnit, IDiscordMessageUnit))

	Start([]*discordgo.ApplicationCommand) error
	Stop()

	GetUser(string) (IDiscordUserUnit, error)
	GetGuild(string) (IDiscordGuildUnit, error)

	BotUser() IDiscordUserUnit
}

type IDiscordCommandFn func(IDiscordInteractionUnit) error

type IDiscordInteractionUnit interface {
	Discord() IDiscordUnit
	Native() *discordgo.InteractionCreate

	DeferReply() error
	Reply(message string) error
	EditReply(message *string) error

	CommandName() string
	IsCommandName(name string) bool
	DispatchEvent(name string, callback IDiscordCommandFn) bool
}

type IDiscordGuildUnit interface {
	Discord() IDiscordUnit
	Native() *discordgo.Guild

	// Base
	Snowflake() string
	Id() string

	// Information
	Name() string
	Description() string
	Icon() string
	Region() string

	IsOwner() bool

	// Methods
	GetChannels() ([]IDiscordChannelUnit, error)
	GetChannel(string) (IDiscordChannelUnit, error)

	GetMemberCount() (int, error)
}

type IDiscordChannelUnit interface {
	Discord() IDiscordUnit
	Native() *discordgo.Channel

	// Base
	Snowflake() string
	Id() string

	// Information
	Guild() (IDiscordGuildUnit, error)
	Name() string
	Topic() string
	Position() int

	NSFW() bool

	Type() discordgo.ChannelType
	Flags() discordgo.ChannelFlags

	// Methods
	FetchMessage(string) (IDiscordMessageUnit, error)
	FetchMessages(limit int) ([]IDiscordMessageUnit, error)
	GetLastMessage() (IDiscordMessageUnit, error)
	SendMessage(string) (IDiscordMessageUnit, error)
	SendMessageOptions(options DiscordMessageSend) (IDiscordMessageUnit, error)

	SendTyping() error
}

type IDiscordMessageUnit interface {
	Discord() IDiscordUnit
	Native() *discordgo.Message

	Channel() IDiscordChannelUnit

	Author() IDiscordUserUnit
	Timestamp() time.Time
	EditedTimestamp() *time.Time
	Mentions() []IDiscordUserUnit

	Edit(message string) error
	EditOptions(options DiscordMessageEdit) error
	Crosspost() error
	Delete() error

	Reply(message string) (IDiscordMessageUnit, error)
}

type IDiscordUserUnit interface {
	Discord() IDiscordUnit
	Native() *discordgo.User

	// Base
	Snowflake() string
	Id() string

	// Information
	Username() string
	Discriminator() string
	GlobalName() string

	IsBot() bool
	IsVerified() bool
	HasMFAEnabled() bool

	IsSystem() bool

	IsAnyNitro() bool
	IsNitroClassic() bool
	IsNitroBasic() bool
	IsNitro() bool
}
