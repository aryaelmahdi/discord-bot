package discord

import (
	"discord-bot/configs"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InitDiscord(config *configs.BotToken) (*discordgo.Session, error) {
	sess, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("cannot create discord session : " + err.Error())
		return nil, err
	}
	fmt.Println("session created")
	return sess, nil
}

func OpenSession(client *discordgo.Session) error {
	if err := client.Open(); err != nil {
		fmt.Println("Cannot open Discord session: " + err.Error())
		return err
	}
	fmt.Println("Session opened")
	return nil
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println("Message received from " + m.Author.Username + ": " + m.Content)
	if m.Content == "hello" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Hello")
		if err != nil {
			fmt.Println("Error sending message: ", err)
		}
	}
}
