package main

import (
	"DiscordBot/config"
	"DiscordBot/handlers"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config.InitConfig()
	dc, err := disgolf.New(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(fmt.Errorf("failed to create bot: %w", err))
	}

	dc.Identify.Intents = discordgo.IntentsAll

	//Init Handlers
	handlers.Handlers(dc)

	//Open Discord Bot
	err = dc.Open()
	if err != nil {
		log.Panic("Failed to opening connection ", err)
	}
	defer dc.Close()

	fmt.Println("Bot esta rodando. Aperte CTRL + C para fechar")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
