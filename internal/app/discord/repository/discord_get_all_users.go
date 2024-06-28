package repository

import "discord-bot/internal/model/domain"

func (repo *DiscordRepositoryImpl) GetAllUsers() []domain.Users {
	var users []domain.Users
	if err := repo.DB.Find(&users).Order("id ASC").Error; err != nil {
		return nil
	}
	return users
}
