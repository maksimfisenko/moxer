package services

import (
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/maksimfisenko/moxer/internal/services/mapper"
	"github.com/maksimfisenko/moxer/internal/services/repo"
)

type authService struct {
	usersRepo repo.UsersRepo
}

func NewAuthSerice(authRepo repo.UsersRepo) *authService {
	return &authService{usersRepo: authRepo}
}

func (as *authService) Register(userDTO *dto.UserDTO) (*dto.UserDTO, error) {
	entity := mapper.FromUserDTOToUserEntity(userDTO)

	if _, err := as.usersRepo.Create(entity); err != nil {
		return nil, err
	}

	return mapper.FromUserEntityToUserDTO(entity), nil
}
