package routes

import (
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

func (routes *DiscordRoutesImpl) InitRoutes(sess *discordgo.Session, cron *cron.Cron) {
	// sess.AddHandler(routes.Handler.CommandHandler)
	sess.AddHandler(routes.Handler.DiscordCommandHandler)
	sess.AddHandler(routes.Handler.AutomatedLogin)
}
