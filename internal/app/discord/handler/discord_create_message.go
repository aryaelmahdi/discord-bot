package handler

import (
	"discord-bot/internal/model/web"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (handler *DiscordHandlerImpl) DiscordCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	fmt.Println(args)
	if args[0] != handler.Prefix {
		return
	}

	fmt.Println("masuk 1")
	if args[1] == "hello" {
		s.ChannelMessageSend(m.ChannelID, "sup bradaah??.")
		return
	}

	if args[1] == "help" {
		s.ChannelMessageSend(m.ChannelID, "toss in 'digidaw login <username> <password> <class>' to do register")
		return
	}

	if args[1] == "register" {
		if len(args) < 5 {
			s.ChannelMessageSend(m.ChannelID, "u stupid or what? username, password and your f ing class are required. ya know what i mean? huh.")
			return
		}
		username := args[2]
		password := args[3]
		class := args[4]
		fmt.Println("username, pass, class", username, password, class)
		s.ChannelMessageSend(m.ChannelID, "wait a sec yea?")
		request := web.DiscordRegister{
			Username: username,
			Password: password,
			Class:    class,
		}
		if err := handler.Service.Register(&request); err != nil {
			if strings.Contains(err.Error(), "exists") {
				s.ChannelMessageSend(m.ChannelID, "ayy dumbass, you are registered in my database so dont call that function again bitch.")
				return
			}
			s.ChannelMessageSend(m.ChannelID, "an error occurred: "+err.Error())
			return
		}
		if err := handler.RunBot6A(username, password); err != nil {
			fmt.Println("error :", err)
			if strings.Contains(err.Error(), "cannot find present button") {
				s.ChannelMessageSend(m.ChannelID, "your status is present already idiot")
				s.ChannelMessageSend(m.ChannelID, "your data is now in my database ^-^.")
				return
			}
			s.ChannelMessageSend(m.ChannelID, "an error occurred: "+err.Error())
			return
		}
		s.ChannelMessageSend(m.ChannelID, "your status is now present. thank me later dog.")
	}
}
