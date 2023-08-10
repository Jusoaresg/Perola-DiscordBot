package models

import "time"

type Phrase struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;type:integer"`
	GuildID   string `gorm:"type:varchar(191)"`
	Creator   string
	Phrase    string
	SavedBy   string
	CreatedAt time.Time
}
