package main

import (
	"discord-bot/configs"
	"discord-bot/internal/app"
	"discord-bot/internal/infrastructure/discord"
	"discord-bot/internal/infrastructure/mysql"
	"discord-bot/internal/infrastructure/selenium"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator"
	"github.com/robfig/cron/v3"
)

func main() {
	config, err := configs.InitConfig()
	if err != nil {
		fmt.Println("cannot initialize config : " + err.Error())
	}

	c := cron.New()
	defer c.Stop()

	validate := validator.New()
	selService, caps, err := selenium.InitSelenium(&config.Selenium)
	if err != nil {
		fmt.Println(err)
	}
	defer selService.Stop()

	db, err := mysql.NewMySQLConnection(&config.MySQL)
	if err != nil {
		fmt.Errorf("cannot connect to mysql : ", err)
	}

	sess, err := discord.InitDiscord(&config.BotToken)
	if err != nil {
		return
	}

	fmt.Println("prefix ", config.BotToken.Prefix)
	app.InitApp(db, validate, sess, config.BotToken.Prefix, &config.Selenium, c, caps)

	c.Start()

	defer sess.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
