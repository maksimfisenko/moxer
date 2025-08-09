package services

import (
	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

type AuthService interface {
	Register(userDTO *dto.User) (*dto.User, error)
	Login(credentials *dto.UserCredentials) (*dto.Token, error)
	GetById(userId uuid.UUID) (*dto.User, error)
}
