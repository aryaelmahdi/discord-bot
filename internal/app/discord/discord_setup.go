package discord

import (
	"discord-bot/configs"
	"discord-bot/internal/app/discord/handler"
	"discord-bot/internal/app/discord/repository"
	"discord-bot/internal/app/discord/routes"
	"discord-bot/internal/app/discord/service"

	"github.com/bwmarrin/discordgo"
	"github.com/go-playground/validator"
)

func DiscordSetup(validator *validator.Validate, sess *discordgo.Session, prefix string, selenium *configs.SeleniumConfig) routes.DiscordRoutes {
	repo := repository.NewDiscordRepository()
	service := service.NewDiscordService(repo, validator)
	handler := handler.NewDiscordHandler(service, prefix, selenium)
	route := routes.NewDiscordRoutes(handler)
	return route
}
