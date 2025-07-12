package entities

import (
	"time"

	"github.com/google/uuid"
)

type Template struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Name      string
	Content   map[string]any `gorm:"type:jsonb;serializer:json"`
	UserId    uuid.UUID      `gorm:"index"`
	User      User           `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
