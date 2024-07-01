package service

import (
	"discord-bot/internal/app/discord/repository"
	"discord-bot/internal/model/domain"
	"discord-bot/internal/model/web"

	"github.com/go-playground/validator"
)

type DiscordService interface {
	Register(request *web.DiscordRegister) error
	Get6A() ([]domain.Users, error)
	Get6B() ([]domain.Users, error)
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
