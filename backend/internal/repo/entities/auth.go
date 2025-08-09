package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `gorm:"primaryKey"`
	Email        string    `gorm:"uniqueIndex"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
