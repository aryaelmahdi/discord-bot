package repository

import (
	"discord-bot/internal/model/domain"

	"gorm.io/gorm"
)

type DiscordRepository interface {
	Register(request *domain.Users) error
	IsExists(username string) bool
	GetAllUsers() []domain.Users
}

type DiscordRepositoryImpl struct {
	DB *gorm.DB
}

func NewDiscordRepository(db *gorm.DB) DiscordRepository {
	return &DiscordRepositoryImpl{
		DB: db,
	}
}
