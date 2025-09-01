package models

import "time"

type URL struct {
	ID          uint   `gorm:"primaryKey"`
	ShortID     string `gorm:"uniqueIndex"`
	OriginalURL string `gorm:"not null"`
	CreatedAt   time.Time
}
