package handler

import (
	"discord-bot/configs"
	"discord-bot/internal/app/discord/service"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

type DiscordHandler interface {
	// Login(s *discordgo.Session, i *discordgo.InteractionCreate) error
	// CommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate)
	DiscordCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate)
	RunBot6A(username, password string) error
	RunBot6B(username, password string) error
	Login6AMonday()
	Login6AMonday2()
	Login6ATuesday()
	Login6AWednesday()
	Login6AWednesday2()
	Login6AThursday()
	Login6AThursday2()
	Login6AFriday()

	Login6BTuesday()
	Login6BTuesday2()
	Login6BWednesday()
	Login6BWednesday2()
	Login6BWednesday3()
	Login6BThursday()
	Login6BThursday2()
	Login6BThursday3()
}

type DiscordHandlerImpl struct {
	Service        service.DiscordService
	Prefix         string
	SeleniumConfig *configs.SeleniumConfig
	Cron           *cron.Cron
}

func NewDiscordHandler(service service.DiscordService, prefix string, selenium *configs.SeleniumConfig, cron *cron.Cron) DiscordHandler {
	return &DiscordHandlerImpl{
		Service:        service,
		Prefix:         prefix,
		SeleniumConfig: selenium,
		Cron:           cron,
	}
}
