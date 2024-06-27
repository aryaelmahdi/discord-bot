package handler

import (
	"discord-bot/configs"
	"discord-bot/internal/app/discord/service"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/tebeka/selenium"
)

type DiscordHandler interface {
	// Login(s *discordgo.Session, i *discordgo.InteractionCreate) error
	// CommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate)
	DiscordCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate)
	RunBot(username, password string) error
	AutomatedLogin6A(s *discordgo.Session, m *discordgo.MessageCreate)
}

type DiscordHandlerImpl struct {
	Service        service.DiscordService
	Prefix         string
	SeleniumConfig *configs.SeleniumConfig
	Cron           *cron.Cron
	Caps           selenium.Capabilities
}

func NewDiscordHandler(service service.DiscordService, prefix string, selenium *configs.SeleniumConfig, cron *cron.Cron, caps selenium.Capabilities) DiscordHandler {
	return &DiscordHandlerImpl{
		Service:        service,
		Prefix:         prefix,
		SeleniumConfig: selenium,
		Cron:           cron,
		Caps:           caps,
	}
}
