package entity

type GuildEntity struct {
	GuildID    string
	ServerName string
	Phrases    []Phrase `gorm:"foreignKey:GuildID;references:GuildID"`
}
