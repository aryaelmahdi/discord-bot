package main

import (
	"discord-bot/configs"
	"discord-bot/internal/app"
	"discord-bot/internal/infrastructure/discord"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator"
)

func main() {
	config, err := configs.InitConfig()
	if err != nil {
		fmt.Println("cannot initialize config : " + err.Error())
	}

	validate := validator.New()

	sess, err := discord.InitDiscord(&config.BotToken)
	if err != nil {
		return
	}

	fmt.Println("prefix ", config.BotToken.Prefix)
	app.InitApp(validate, sess, config.BotToken.Prefix)

	defer sess.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
