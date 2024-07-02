package app

import (
	"discord-bot/configs"
	discordRoutesPkg "discord-bot/internal/app/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/go-playground/validator"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB, validator *validator.Validate, sess *discordgo.Session, prefix string, selenium *configs.SeleniumConfig, cron *cron.Cron) {
	discordRoute := discordRoutesPkg.DiscordSetup(db, validator, sess, prefix, selenium, cron)
	discordRoute.InitRoutes(sess)
}
