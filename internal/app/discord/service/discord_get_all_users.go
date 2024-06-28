package service

import (
	"discord-bot/internal/model/domain"
	"fmt"
)

func (service *DiscordServiceImpl) GetAllUsers() ([]domain.Users, error) {
	res := service.Repository.GetAllUsers()
	if len(res) == 0 {
		return nil, fmt.Errorf("no users found")
	}
	return res, nil
}
