package web

type DiscordRegister struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
