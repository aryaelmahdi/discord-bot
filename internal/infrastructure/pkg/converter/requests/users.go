package requests

import (
	"discord-bot/internal/model/domain"
	"discord-bot/internal/model/web"
	"fmt"
)

func RegisterRegisterToDomain(request *web.DiscordRegister) (*domain.Users, error) {
	if request.Class != "CCIT6A" && request.Class != "CCIT6B" {
		return nil, fmt.Errorf("class not valid")
	}
	return &domain.Users{
		Username: request.Username,
		Password: request.Password,
		Class:    request.Class,
	}, nil
}
