package service

import (
	"discord-bot/internal/model/domain"
	"fmt"
)

func (service *DiscordServiceImpl) Get6B() ([]domain.Users, error) {
	res := service.Repository.Get6B()
	if len(res) == 0 {
		return nil, fmt.Errorf("no users found")
	}
	return res, nil
}
