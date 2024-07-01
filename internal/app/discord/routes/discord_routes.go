package routes

import (
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

func (routes *DiscordRoutesImpl) InitRoutes(sess *discordgo.Session, cron *cron.Cron) {
	// sess.AddHandler(routes.Handler.CommandHandler)
	sess.AddHandler(routes.Handler.DiscordCommandHandler)

	// Login 6A
	sess.AddHandler(routes.Handler.Login6AMonday)
	sess.AddHandler(routes.Handler.Login6AMonday2)
	sess.AddHandler(routes.Handler.Login6ATuesday)
	sess.AddHandler(routes.Handler.Login6AWednesday)
	sess.AddHandler(routes.Handler.Login6AWednesday2)
	sess.AddHandler(routes.Handler.Login6AThursday)
	sess.AddHandler(routes.Handler.Login6AThursday2)
	sess.AddHandler(routes.Handler.Login6AFriday)

	// Login 6B
	sess.AddHandler(routes.Handler.Login6BTuesday)
	sess.AddHandler(routes.Handler.Login6BTuesday2)
	sess.AddHandler(routes.Handler.Login6BWednesday)
	sess.AddHandler(routes.Handler.Login6BWednesday2)
	sess.AddHandler(routes.Handler.Login6BWednesday3)
	sess.AddHandler(routes.Handler.Login6BThursday)
	sess.AddHandler(routes.Handler.Login6BThursday2)
	sess.AddHandler(routes.Handler.Login6BThursday3)
}
