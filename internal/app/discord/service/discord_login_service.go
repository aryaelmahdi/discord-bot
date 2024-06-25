package service

import (
	"discord-bot/internal/model/web"
)

func (service *DiscordServiceImpl) Login(request *web.DiscordLogin) error {
	if err := service.Validate.Struct(request); err != nil {
		return err
	}
	if err := service.Repository.Login(request.Username, request.Password); err != nil {
		return err
	}
	return nil
}
