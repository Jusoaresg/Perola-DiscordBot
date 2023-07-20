package db

import (
	"DiscordBot/infra/entity"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// AutoMigrate caso necessario
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("infra/db/sql.db"), &gorm.Config{})
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
	db, err := gorm.Open(sqlite.Open("sql.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("error on open database")
	}
	err = db.AutoMigrate(&entity.GuildEntity{}, &entity.Phrase{})
	if err != nil {
		return nil, errors.New("error on migrate")
	}
	return db, nil
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
