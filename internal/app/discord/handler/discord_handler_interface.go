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
	Login6AMonday(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6AMonday2(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6ATuesday(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6AWednesday(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6AWednesday2(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6AThursday(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6AThursday2(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6AFriday(s *discordgo.Session, m *discordgo.MessageCreate)

	Login6BTuesday(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6BTuesday2(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6BWednesday(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6BWednesday2(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6BWednesday3(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6BThursday(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6BThursday2(s *discordgo.Session, m *discordgo.MessageCreate)
	Login6BThursday3(s *discordgo.Session, m *discordgo.MessageCreate)
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
