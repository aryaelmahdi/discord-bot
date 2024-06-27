package discord

import (
	"discord-bot/configs"
	"discord-bot/internal/app/discord/handler"
	"discord-bot/internal/app/discord/repository"
	"discord-bot/internal/app/discord/routes"
	"discord-bot/internal/app/discord/service"

	"github.com/bwmarrin/discordgo"
	"github.com/go-playground/validator"
	"github.com/robfig/cron/v3"
	"github.com/tebeka/selenium"
)

func DiscordSetup(validator *validator.Validate, sess *discordgo.Session, prefix string, selenium *configs.SeleniumConfig, cron *cron.Cron, caps selenium.Capabilities) routes.DiscordRoutes {
	repo := repository.NewDiscordRepository()
	service := service.NewDiscordService(repo, validator)
	handler := handler.NewDiscordHandler(service, prefix, selenium, cron, caps)
	route := routes.NewDiscordRoutes(handler)
	return route
}
