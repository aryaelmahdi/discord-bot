package repository

type DiscordRepository interface {
	Login(username string, password string) error
}

type DiscordRepositoryImpl struct{}

func NewDiscordRepository() DiscordRepository {
	return &DiscordRepositoryImpl{}
}
