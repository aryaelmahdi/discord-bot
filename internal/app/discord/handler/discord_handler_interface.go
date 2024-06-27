package handler

import (
	"discord-bot/configs"
	"discord-bot/internal/app/discord/service"

	"github.com/bwmarrin/discordgo"
)

type DiscordHandler interface {
	Login(s *discordgo.Session, i *discordgo.InteractionCreate) error
	// CommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate)
	MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
	RunBot(username, password string) error
}

type DiscordHandlerImpl struct {
	Service        service.DiscordService
	Prefix         string
	SeleniumConfig *configs.SeleniumConfig
}

func NewDiscordHandler(service service.DiscordService, prefix string, selenium *configs.SeleniumConfig) DiscordHandler {
	return &DiscordHandlerImpl{
		Service:        service,
		Prefix:         prefix,
		SeleniumConfig: selenium,
	}
}
