package web

type DiscordLogin struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
