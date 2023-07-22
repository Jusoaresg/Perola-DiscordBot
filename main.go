package main

import (
	"DiscordBot/infra/handler"
	"fmt"
	"github.com/FedorLap2006/disgolf"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

// INIT: init env variables
func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic("error on get wd", err)
	}
	envPath := filepath.Join(cwd, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Panic(fmt.Errorf("error on loading .env: %w", err))
	}
}

func main() {
	dc, err := disgolf.New(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(fmt.Errorf("failed to create bot: %w", err))
	}

	dc.Identify.Intents = discordgo.IntentsAll

	//Init Handlers
	handler.Handlers(dc)

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
