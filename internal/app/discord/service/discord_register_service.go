package service

import (
	"discord-bot/internal/infrastructure/pkg/converter/requests"
	"discord-bot/internal/model/web"
	"fmt"
)

func (service *DiscordServiceImpl) Register(request *web.DiscordRegister) error {
	if err := service.Validate.Struct(request); err != nil {
		return err
	}
	if exists := service.Repository.IsExists(request.Username); exists {
		return fmt.Errorf("user exists")
	}
	domain, err := requests.RegisterRegisterToDomain(request)
	if err != nil {
		return err
	}
	if err := service.Repository.Register(domain); err != nil {
		return err
	}
	return nil
}
