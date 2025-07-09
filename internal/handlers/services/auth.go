package services

import "github.com/maksimfisenko/moxer/internal/services/dto"

type AuthService interface {
	Register(userDTO *dto.UserDTO) (*dto.UserDTO, error)
}
