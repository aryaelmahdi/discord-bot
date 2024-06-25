package service

import (
	"discord-bot/internal/app/discord/repository"
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

func NewDiscordService(repo repository.DiscordRepository, validator *validator.Validate) DiscordService {
	return &DiscordServiceImpl{
		Repository: repo,
		Validate:   validator,
	}
}
