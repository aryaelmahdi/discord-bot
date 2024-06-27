package app

import (
	"discord-bot/configs"
	discordRoutesPkg "discord-bot/internal/app/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/go-playground/validator"
)

func InitApp(validator *validator.Validate, sess *discordgo.Session, prefix string, selenium *configs.SeleniumConfig) {
	discordRoute := discordRoutesPkg.DiscordSetup(validator, sess, prefix, selenium)
	discordRoute.InitRoutes(sess)
}
