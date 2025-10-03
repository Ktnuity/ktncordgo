package ktncordgo

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

// CreateDiscordUnit takes a discord token and creates a [DiscordUnit] instance.
//
// Returns the create instance on success, otherwise an error.
func CreateDiscordUnit(token string) (IDiscordUnit, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("failed to create discord session: %w", err)
	}

	session.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	return &DiscordUnit{session: session}, nil
}

// Start takes a slice of commands, opens the session, and registers the commands with discord.
//
// Parameters:
//   commands - a slice of [discordgo.ApplicationCommand] references.
//
// Returns an error on failure.
func (self *DiscordUnit) Start(commands []*discordgo.ApplicationCommand) error {
	err := self.session.Open()
	if err != nil {
		return fmt.Errorf("failed to open session: %w", err)
	}

	for cmd := range commands {
		_, err = self.session.ApplicationCommandCreate(self.session.State.User.ID, "", commands[cmd])

		if err != nil {
			log.Printf("failed to create slash command: %v", err)
		}
	}

	return nil
}

// Stop stops the discord session.
func (self *DiscordUnit) Stop() {
	self.session.Close()
}

// NewDiscordUnit takes a [discordgo.Session] reference and creates a [DiscordUnit].
//
// Parameters:
//   session - The underlying [discordgo.Session] reference to use.
//
// Returns the created [DiscordUnit] reference.
func NewDiscordUnit(session *discordgo.Session) IDiscordUnit {
	return &DiscordUnit{
		session: session,
	}
}

// Session returns a reference to the underlying [discordgo.Session] object.
//
// Returns the [discordgo.Session] reference.
func (self *DiscordUnit) Session() *discordgo.Session {
	return self.session
}

// NewInteractionUnit creates a new [DiscordInteractionUnit] reference using the current [DiscordUnit] instance as the parent object.
//
// Parameters:
//   interaction - The underlying [discordgo.InteractionCreate] instance to use.
//
// Returns the created [DiscordInteractionUnit] reference.
//
// See: [DiscordInteractionUnit]
func (self *DiscordUnit) NewInteractionUnit(interaction *discordgo.InteractionCreate) IDiscordInteractionUnit {
	return &DiscordInteractionUnit {
		discord: self,
		interaction: interaction,
	}
}

// OnSlashCommand registers an event handler for Slash Commands.
//
// Parameters:
//   callback - The callback handler for the slash command event.
func (self *DiscordUnit) OnSlashCommand(callback func(IDiscordUnit, IDiscordInteractionUnit)) {
	self.session.AddHandler(func (inSession *discordgo.Session, inInteraction *discordgo.InteractionCreate) {
		session := NewDiscordUnit(inSession)
		interaction := session.NewInteractionUnit(inInteraction)

		callback(session, interaction)
	})
}

// OnMessageCreate registers an event handler for Channel Chat Messages.
//
// Parameters:
//   callback - The callback handler for the message create event.
func (self *DiscordUnit) OnMessageCreate(callback func (IDiscordUnit, IDiscordMessageUnit)) {
	self.session.AddHandler(func (inSession *discordgo.Session, inMessage *discordgo.MessageCreate) {
		var session *DiscordUnit = &DiscordUnit{session: inSession}
		var message *DiscordMessageUnit = &DiscordMessageUnit{
			discord: session,
			message: inMessage.Message,
		}

		callback(session, message)
	})
}

// GetUser finds and returns a user given a snowflake ID.
//
// Parameters:
//   userId - The ID of the user to find.
//
// Returns the user object if found, otherwise an error.
// 
// See: [DiscordUserUnit]
func (self *DiscordUnit) GetUser(userId string) (IDiscordUserUnit, error) {
	user, err := self.session.User(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch discord user: %v", err)
	}

	return &DiscordUserUnit{
		discord: self,
		user: user,
	}, nil
}

// GetChannel finds and returns a channel given a snowflake ID.
//
// Parameters:
//   channelId - The ID of the channel to find.
//
// Returns the channel object if found, otherwise an error.
//
// See: [DiscordChannelUnit]
func (self *DiscordUnit) GetChannel(channelId string) (IDiscordChannelUnit, error) {
	channel, err := self.session.Channel(channelId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch discord channel: %v", err)
	}

	return &DiscordChannelUnit{
		discord: self,
		channel: channel,
	}, nil
}

// GetGuild finds and returns a guild given a snowflake ID.
//
// Parameters:
//   guildId - The ID of the guild to find.
//
// Returns the guild object if found, otherwise an error.
//
// See: [DiscordGuildUnit]
func (self *DiscordUnit) GetGuild(guildId string) (IDiscordGuildUnit, error) {
	guild, err := self.session.Guild(guildId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch discord guild: %v", err)
	}

	return &DiscordGuildUnit{
		discord: self,
		guild: guild,
	}, nil
}

// BotUser gets the underlying bot user of the [DiscordUnit] instance.
//
// Returns the user instance of present, otherwise nil.
//
// See: [DiscordUserUnit]
func (self *DiscordUnit) BotUser() IDiscordUserUnit {
	user := self.session.State.User

	if user == nil {
		return nil
	}

	return &DiscordUserUnit{
		discord: self,
		user: user,
	}
}

// Snowflake gets the ID of the discord unit's bot user.
//
// See: [DiscordUserUnit.Snowflake]
func (self *DiscordUnit) BotSnowflake() string {
	return self.BotUser().Snowflake()
}

// Id gets the ID of the discord unit's bot user
//
// See: [DiscordUnit.Snowflake]
// See: [DiscordUserUnit.Id]
func (self *DiscordUnit) BotId() string {
	return self.BotUser().Id()
}
