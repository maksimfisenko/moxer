package dto

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Credentials struct {
	Email    string
	Password string
}

type Token struct {
	Token string
}
