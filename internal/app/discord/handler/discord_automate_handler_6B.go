package handler

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (handler *DiscordHandlerImpl) Login6BTuesday(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("30 7 * * 1", func() {
		fmt.Println("test job 6B")
		res, err := handler.Service.Get6B()
		if err != nil {
			fmt.Errorf(err.Error())
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		fmt.Println("res B : ", res)
		for _, user := range res {
			fmt.Println("user : ", user.Username)
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

func (handler *DiscordHandlerImpl) Login6BTuesday2(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("50 10 * * 1", func() {
		fmt.Println("monday job 2")
		res, err := handler.Service.Get6B()
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
}

func (handler *DiscordHandlerImpl) Login6BWednesday(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("30 7 * * 2", func() {
		fmt.Println("tuesday job")
		res, err := handler.Service.Get6B()
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
}

func (handler *DiscordHandlerImpl) Login6BWednesday2(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("50 10 * * 3", func() {
		fmt.Println("monday job 2")
		res, err := handler.Service.Get6B()
		if err != nil {
			fmt.Errorf(err.Error())
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot(user.Username, user.Password); err != nil {
				fmt.Errorf(err.Error())
				time.Sleep(3 * time.Second)
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

func (handler *DiscordHandlerImpl) Login6BWednesday3(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("15 15 * * 3", func() {
		res, err := handler.Service.Get6B()
		if err != nil {
			fmt.Errorf(err.Error())
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot(user.Username, user.Password); err != nil {
				fmt.Errorf(err.Error())
				time.Sleep(3 * time.Second)
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

func (handler *DiscordHandlerImpl) Login6BThursday(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("30 7 * * 4", func() {
		res, err := handler.Service.Get6B()
		if err != nil {
			fmt.Errorf(err.Error())
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot(user.Username, user.Password); err != nil {
				fmt.Errorf(err.Error())
				time.Sleep(3 * time.Second)
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

func (handler *DiscordHandlerImpl) Login6BThursday2(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("50 10 * * 4", func() {
		res, err := handler.Service.Get6B()
		if err != nil {
			fmt.Errorf(err.Error())
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot(user.Username, user.Password); err != nil {
				fmt.Errorf(err.Error())
				time.Sleep(3 * time.Second)
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

func (handler *DiscordHandlerImpl) Login6BThursday3(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("15 15 * * 5", func() {
		res, err := handler.Service.Get6B()
		if err != nil {
			fmt.Errorf(err.Error())
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		for _, user := range res {
			fmt.Println("user : ", user.Username)
			if err := handler.RunBot(user.Username, user.Password); err != nil {
				fmt.Errorf(err.Error())
				time.Sleep(3 * time.Second)
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
