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
		return
	}

	if args[1] == "hello" {
		fmt.Println("args hello")
		s.ChannelMessageSend(m.ChannelID, "fuck you bitch.")
		return
	}

	if args[1] == "help" {
		fmt.Println("args hello")
		s.ChannelMessageSend(m.ChannelID, "Use 'digidaw login <username> <password>' to do login")
		return
	}
	username, password := "", ""
	if args[1] == "login" {
		fmt.Println("args login ", len(args))
		if len(args) < 4 {
			s.ChannelMessageSend(m.ChannelID, "u stupid or what? username and password are required")
			return
		}
		if args[2] == "" || args[3] == "" {
			s.ChannelMessageSend(m.ChannelID, "u stupid or what? username and password are required")
			return
		}
		username = args[2]
		password = args[3]
		fmt.Println("username & pass", username, password)
		s.ChannelMessageSend(m.ChannelID, "wait a sec yea?")
	}
	if err := handler.RunBot(username, password); err != nil {
		fmt.Println("error :", err)
		if strings.Contains(err.Error(), "cannot find present button") {
			s.ChannelMessageSend(m.ChannelID, "your status is present already idiot")
			return
		}
		return
	}
	s.ChannelMessageSend(m.ChannelID, "yeay logged in to academia pnj")
}
