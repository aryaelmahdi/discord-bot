package repository

import "discord-bot/internal/model/domain"

func (repo *DiscordRepositoryImpl) IsExists(username string) bool {
	var user domain.Users
	if err := repo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return false
	}
	return true
}
