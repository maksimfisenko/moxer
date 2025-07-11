package services

import (
	"errors"

	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/maksimfisenko/moxer/internal/services/jwt"
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

func (as *authService) Login(credentials *dto.UserCredentials) (*dto.Token, error) {
	user, err := as.usersRepo.FindByEmail(credentials.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != credentials.Password {
		return nil, errors.New("user not found")
	}

	token, err := jwt.GenerateToken(user.Id.String())
	if err != nil {
		return nil, err
	}

	return &dto.Token{Token: token}, nil
}
