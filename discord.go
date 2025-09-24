package ktncordgo

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func CreateDiscordUnit(token string) (IDiscordUnit, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("failed to create discord session: %w", err)
	}

	session.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	return &DiscordUnit{session: session}, nil
}

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

func (self *DiscordUnit) Stop() {
	self.session.Close()
}

func NewDiscordUnit(session *discordgo.Session) IDiscordUnit {
	return &DiscordUnit{
		session: session,
	}
}

func (self *DiscordUnit) Session() *discordgo.Session {
	return self.session
}

func (self *DiscordUnit) NewInteractionUnit(interaction *discordgo.InteractionCreate) IDiscordInteractionUnit {
	return &DiscordInteractionUnit {
		discord: self,
		interaction: interaction,
	}
}

func (self *DiscordUnit) OnSlashCommand(callback func(IDiscordUnit, IDiscordInteractionUnit)) {
	self.session.AddHandler(func (inSession *discordgo.Session, inInteraction *discordgo.InteractionCreate) {
		session := NewDiscordUnit(inSession)
		interaction := session.NewInteractionUnit(inInteraction)

		callback(session, interaction)
	})
}

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
