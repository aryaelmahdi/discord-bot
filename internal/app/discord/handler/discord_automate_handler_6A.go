package handler

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (handler *DiscordHandlerImpl) Login6AMonday(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("masuk automated login")
	_, err := handler.Cron.AddFunc("50 10 * * 1", func() {
		fmt.Println("monday job")
		res, err := handler.Service.Get6A()
		if err != nil {
			fmt.Errorf(err.Error())
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		fmt.Println("res A : ", res)
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

func (handler *DiscordHandlerImpl) Login6AMonday2(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("15 15 * * 1", func() {
		fmt.Println("monday job 2")
		res, err := handler.Service.Get6A()
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

func (handler *DiscordHandlerImpl) Login6ATuesday(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("30 7 * * 2", func() {
		fmt.Println("tuesday job")
		res, err := handler.Service.Get6A()
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

func (handler *DiscordHandlerImpl) Login6AWednesday(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("30 7 * * 3", func() {
		fmt.Println("monday job 2")
		res, err := handler.Service.Get6A()
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

func (handler *DiscordHandlerImpl) Login6AWednesday2(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("50 10 * * 3", func() {
		res, err := handler.Service.Get6A()
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

func (handler *DiscordHandlerImpl) Login6AThursday(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("30 7 * * 4", func() {
		res, err := handler.Service.Get6A()
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

func (handler *DiscordHandlerImpl) Login6AThursday2(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("15 15 * * 4", func() {
		res, err := handler.Service.Get6A()
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

func (handler *DiscordHandlerImpl) Login6AFriday(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := handler.Cron.AddFunc("20 13 * * 5", func() {
		res, err := handler.Service.Get6A()
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
