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

	if args[1] == "register" {
		fmt.Println("args register ", len(args))
		if len(args) < 5 {
			s.ChannelMessageSend(m.ChannelID, "u stupid or what? username, password and your f ing class are required")
			return
		}
		username := args[2]
		password := args[3]
		class := args[4]
		fmt.Println("username, pass, class", username, password, class)
		s.ChannelMessageSend(m.ChannelID, "wait a sec yea?")
		// cron job
		// _, err := handler.Cron.AddFunc("* * * * *", func() {
		// 	if err := handler.RunBot(username, password); err != nil {
		// 		fmt.Println("error :", err)
		// 		if strings.Contains(err.Error(), "cannot find present button") {
		// 			s.ChannelMessageSend(m.ChannelID, "your status is present already idiot")
		// 			return
		// 		}
		// 		s.ChannelMessageSend(m.ChannelID, "an error occurred: "+err.Error())
		// 		return
		// 	}
		// 	s.ChannelMessageSend(m.ChannelID, "your status is now present. thank me later dog.")
		// })
		// if err != nil {
		// 	log.Fatalf("Error scheduling cron job: %v", err)
		// }
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
		if err := handler.RunBot(username, password); err != nil {
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
