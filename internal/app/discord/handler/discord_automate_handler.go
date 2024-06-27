package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (handler *DiscordHandlerImpl) AutomatedLogin(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("masuk automated login")
	_, err := handler.Cron.AddFunc("* * * * *", func() {
		res, err := handler.Service.GetAllUsers()
		if err != nil {
			fmt.Errorf(err.Error())
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		for _, user := range res {
			if err := handler.RunBot(user.Username, user.Password); err != nil {
				fmt.Errorf(err.Error())
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("hey, i aint gonna lie i cant logged in to %s", user.Username))
				return
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("thank me later %s for this shit. youre in", user.Username))
			return
		}
	})
	if err != nil {
		fmt.Errorf("cron job error: ", err)
	}
}
