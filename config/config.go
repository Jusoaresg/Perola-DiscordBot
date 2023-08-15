package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func InitConfig() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic("error on get wd", err)
	}
	// ,/config/bot.env
	envPath := filepath.Join(cwd, "config/bot.env")

	err = godotenv.Load(envPath)
	if err != nil {
		log.Panic(fmt.Errorf("error on loading .env: %w", err))
	}
}
