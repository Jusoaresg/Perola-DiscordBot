package entity

import "time"

type Phrase struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;type:integer"`
	GuildID   string //`gorm:"index"`
	Creator   string
	Phrase    string
	SavedBy   string
	CreatedAt time.Time
}
