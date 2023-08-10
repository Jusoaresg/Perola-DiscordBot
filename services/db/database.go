package db

import (
	entity "DiscordBot/models"
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// AutoMigrate caso necessario
func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATA_BASE_DSN")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		return nil, errors.New("error on open database")
	}
	err = db.AutoMigrate(&entity.GuildEntity{}, &entity.Phrase{})
	if err != nil {
		return nil, errors.New("error on migrate")
	}
	return db, nil
}

func OpenDb() (*gorm.DB, error) {
	dsn := os.Getenv("DATA_BASE_DSN")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		return nil, errors.New("error on open database")
	}
	return db, nil
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()
	if err != nil {
		return errors.New("Error on get sql db")
	}
	err = database.Close()
	if err != nil {
		return errors.New("Error on close db")
	}
	return nil
}

func SearchIfGuildExist(db *gorm.DB, guildId string, serverName string) (*gorm.DB, error) {
	//Detect if the column guild id exist
	listServers := []entity.GuildEntity{}
	db.Find(&listServers)
	for _, sv := range listServers {
		if sv.GuildID == guildId {
			return db, nil
		}
	}

	//Creating a new Server Table
	server := entity.GuildEntity{
		GuildID:    guildId,
		ServerName: serverName,
		Phrases:    []entity.Phrase{},
	}
	db.Create(&server)
	return db, nil
}
