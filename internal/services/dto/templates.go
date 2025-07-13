package dto

import (
	"time"

	"github.com/google/uuid"
)

type Template struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Name      string
	Content   map[string]any
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GeneratedData struct {
	Data []map[string]any
}
