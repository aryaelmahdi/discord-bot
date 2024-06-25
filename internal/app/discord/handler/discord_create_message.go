package handler

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (handler *DiscordHandlerImpl) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println("Message received from " + m.Author.Username + ": " + m.Content)
	if m.Content == "hello" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Hello")
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}

	args := strings.Split(m.Content, " ")
	if args[0] != handler.Prefix {
		fmt.Println("no prefix")
		return
	}

	if args[1] == "hello" {
		fmt.Println("prefix args 1")
		s.ChannelMessageSend(m.ChannelID, "fuck you bitch.")
	}
}
