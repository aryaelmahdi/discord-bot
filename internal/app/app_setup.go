package app

import (
	discordRoutesPkg "discord-bot/internal/app/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/go-playground/validator"
)

func InitApp(validator *validator.Validate, sess *discordgo.Session, prefix string) {
	discordRoute := discordRoutesPkg.DiscordSetup(validator, sess, prefix)
	discordRoute.InitRoutes(sess)
}
