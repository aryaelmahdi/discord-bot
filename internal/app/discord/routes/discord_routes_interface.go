package routes

import (
	"discord-bot/internal/app/discord/handler"

	"github.com/bwmarrin/discordgo"
)

type DiscordRoutes interface {
	InitRoutes(sess *discordgo.Session)
}

type DiscordRoutesImpl struct {
	Handler handler.DiscordHandler
}

func NewDiscordRoutes(handler handler.DiscordHandler) DiscordRoutes {
	return &DiscordRoutesImpl{
		Handler: handler,
	}
}
