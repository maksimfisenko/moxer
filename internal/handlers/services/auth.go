package services

import (
	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

type AuthService interface {
	Register(userDTO *dto.UserDTO) (*dto.UserDTO, error)
	Login(credentials *dto.UserCredentials) (*dto.Token, error)
	GetById(userId uuid.UUID) (*dto.UserDTO, error)
}
