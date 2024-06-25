package handler

import (
	"discord-bot/internal/app/discord_commands/service"

	"github.com/bwmarrin/discordgo"
)

type DiscordHandler interface {
	Login(s *discordgo.Session, i *discordgo.InteractionCreate) error
}

type DiscordHandlerImpl struct {
	Service service.DiscordService
}

func NewDiscordHandler(service service.DiscordService) DiscordHandler {
	return &DiscordHandlerImpl{
		Service: service,
	}
}
