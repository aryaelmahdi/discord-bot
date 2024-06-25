package service

import (
	"discord-bot/internal/app/discord_commands/repository"
	"discord-bot/internal/model/web"

	"github.com/go-playground/validator"
)

type DiscordService interface {
	Login(request *web.DiscordLogin) error
}

type DiscordServiceImpl struct {
	Repository repository.DiscordRepository
	Validate   *validator.Validate
}

func NewDiscordService(validator *validator.Validate) DiscordService {
	return &DiscordServiceImpl{
		Validate: validator,
	}
}
