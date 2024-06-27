package repository

import (
	"discord-bot/internal/model/domain"
)

func (repo *DiscordRepositoryImpl) Register(request *domain.Users) error {
	if err := repo.DB.Create(&request).Error; err != nil {
		return err
	}
	return nil
}
