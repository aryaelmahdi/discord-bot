package main

import (
	"discord-bot/configs"
	"discord-bot/internal/infrastructure/discord"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config, err := configs.InitConfig()
	if err != nil {
		fmt.Println("cannot initialize config : " + err.Error())
	}

	sess, err := discord.InitDiscord(&config.BotToken)
	if err != nil {
		return
	}

	sess.AddHandler(discord.MessageCreate)
	fmt.Println("session handler added")
	defer sess.Close()

	err = discord.OpenSession(sess)
	if err != nil {
		return
	}

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
