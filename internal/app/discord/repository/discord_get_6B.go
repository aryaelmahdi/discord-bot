package repository

import "discord-bot/internal/model/domain"

func (repo *DiscordRepositoryImpl) Get6B() []domain.Users {
	var users []domain.Users
	if err := repo.DB.Where("class = ?", "CCIT6B").Find(&users).Order("id ASC").Error; err != nil {
		return nil
	}
	return users
}
