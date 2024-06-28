package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (handler *DiscordHandlerImpl) AutomatedLogin6A(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("masuk automated login")
	// monday
	_, err := handler.Cron.AddFunc("50 10 * * 1", func() {
		fmt.Println("monday job")
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
				continue
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("thank me later %s for this shit. youre in", user.Username))
		}
	})
	if err != nil {
		fmt.Errorf("cron job error: ", err)
	}

	_, err = handler.Cron.AddFunc("15 15 * * 1", func() {
		fmt.Println("monday job 2")
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
				continue
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("thank me later %s for this shit. youre in", user.Username))
		}
	})
	if err != nil {
		fmt.Errorf("cron job error: ", err)
	}

	//tuesday
	_, err = handler.Cron.AddFunc("30 7 * * 2", func() {
		fmt.Println("tuesday job")
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
				continue
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("thank me later %s for this shit. youre in", user.Username))
		}
	})
	if err != nil {
		fmt.Errorf("cron job error: ", err)
	}

	//wednesday
	_, err = handler.Cron.AddFunc("30 7 * * 3", func() {
		fmt.Println("wednesday job")
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
				continue
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("thank me later %s for this shit. youre in", user.Username))
		}
	})
	if err != nil {
		fmt.Errorf("cron job error: ", err)
	}

	_, err = handler.Cron.AddFunc("50 10 * * 3", func() {
		fmt.Println("wednesday job 2")
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
				continue
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("thank me later %s for this shit. youre in", user.Username))
		}
	})
	if err != nil {
		fmt.Errorf("cron job error: ", err)
	}

	//thursday
	_, err = handler.Cron.AddFunc("30 7 * * 4", func() {
		fmt.Println("thursday job")
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
				continue
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("thank me later %s for this shit. youre in", user.Username))
		}
	})
	if err != nil {
		fmt.Errorf("cron job error: ", err)
	}

	_, err = handler.Cron.AddFunc("15 15 * * 4", func() {
		fmt.Println("thursday job 2")
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
				continue
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("thank me later %s for this shit. youre in", user.Username))
		}
	})
	if err != nil {
		fmt.Errorf("cron job error: ", err)
	}

	//friday
	_, err = handler.Cron.AddFunc("20 13 * * 5", func() {
		fmt.Println("friday job")
		res, err := handler.Service.GetAllUsers()
		fmt.Println(res)
		if err != nil {
			fmt.Errorf(err.Error())
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user :", user)
			if err := handler.RunBot(user.Username, user.Password); err != nil {
				fmt.Errorf(err.Error())
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("hey, i aint gonna lie i cant logged in to %s", user.Username))
				continue
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("thank me later %s for this shit. youre in", user.Username))
		}
	})
	if err != nil {
		fmt.Errorf("cron job error: ", err)
	}
}
