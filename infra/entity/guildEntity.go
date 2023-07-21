package entity

type GuildEntity struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	GuildID    string `gorm:"type:varchar(191);unique"`
	ServerName string
	Phrases    []Phrase `gorm:"foreignKey:GuildID;references:GuildID"`
}
