package app

import (
	"discord-bot/configs"
	discordRoutesPkg "discord-bot/internal/app/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/go-playground/validator"
	"github.com/robfig/cron/v3"
	"github.com/tebeka/selenium"
)

func InitApp(validator *validator.Validate, sess *discordgo.Session, prefix string, selenium *configs.SeleniumConfig, cron *cron.Cron, caps selenium.Capabilities) {
	discordRoute := discordRoutesPkg.DiscordSetup(validator, sess, prefix, selenium, cron, caps)
	discordRoute.InitRoutes(sess)
}
