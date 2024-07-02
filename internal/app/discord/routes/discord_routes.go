package routes

import (
	"github.com/bwmarrin/discordgo"
)

func (routes *DiscordRoutesImpl) InitRoutes(sess *discordgo.Session) {
	// sess.AddHandler(routes.Handler.CommandHandler)
	sess.AddHandler(routes.Handler.DiscordCommandHandler)

	// Login 6A
	go func() {
		routes.Handler.Login6AMonday()
		routes.Handler.Login6AMonday2()
		routes.Handler.Login6ATuesday()
		routes.Handler.Login6AWednesday()
		routes.Handler.Login6AWednesday2()
		routes.Handler.Login6AThursday()
		routes.Handler.Login6AThursday2()
		routes.Handler.Login6AFriday()
	}()

	// Login 6B
	go func() {
		routes.Handler.Login6BTuesday()
		routes.Handler.Login6BTuesday2()
		routes.Handler.Login6BWednesday()
		routes.Handler.Login6BWednesday2()
		routes.Handler.Login6BWednesday3()
		routes.Handler.Login6BThursday()
		routes.Handler.Login6BThursday2()
		routes.Handler.Login6BThursday3()
	}()
}
