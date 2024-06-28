package discord

// import (
// 	"fmt"

// 	"github.com/bwmarrin/discordgo"
// )

// func RegisterCommands(s *discordgo.Session) {
// 	commands := []*discordgo.ApplicationCommand{
// 		{
// 			Name:        "ping",
// 			Description: "Nothing special, just ping ya know...'",
// 		},
// 	}

// 	for _, v := range commands {
// 		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
// 		if err != nil {
// 			fmt.Printf("Cannot create '%v' command: %v\n", v.Name, err)
// 		}
// 	}
// }
