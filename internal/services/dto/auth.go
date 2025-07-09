package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserDTO struct {
	Id        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
