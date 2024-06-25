package routes

import "github.com/bwmarrin/discordgo"

func (routes *DiscordRoutesImpl) InitRoutes(sess *discordgo.Session) {
	// sess.AddHandler(routes.Handler.CommandHandler)
	sess.AddHandler(routes.Handler.MessageCreate)
}
