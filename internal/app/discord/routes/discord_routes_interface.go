package routes

import (
	"discord-bot/internal/app/discord/handler"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

type DiscordRoutes interface {
	InitRoutes(sess *discordgo.Session, cron *cron.Cron)
}

type DiscordRoutesImpl struct {
	Handler handler.DiscordHandler
}

func NewDiscordRoutes(handler handler.DiscordHandler) DiscordRoutes {
	return &DiscordRoutesImpl{
		Handler: handler,
	}
}
