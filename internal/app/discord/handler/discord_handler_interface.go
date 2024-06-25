package handler

import (
	"discord-bot/internal/app/discord/service"

	"github.com/bwmarrin/discordgo"
)

type DiscordHandler interface {
	Login(s *discordgo.Session, i *discordgo.InteractionCreate) error
	// CommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate)
	MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)
}

type DiscordHandlerImpl struct {
	Service service.DiscordService
	Prefix  string
}

func NewDiscordHandler(service service.DiscordService, prefix string) DiscordHandler {
	return &DiscordHandlerImpl{
		Service: service,
		Prefix:  prefix,
	}
}
