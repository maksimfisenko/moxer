package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/errorsx"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/maksimfisenko/moxer/internal/services/jwt"
	"github.com/maksimfisenko/moxer/internal/services/mapper"
	"github.com/maksimfisenko/moxer/internal/services/repo"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	usersRepo repo.UsersRepo
}

func NewAuthSerice(authRepo repo.UsersRepo) *authService {
	return &authService{usersRepo: authRepo}
}

func (as *authService) Register(userDTO *dto.User) (*dto.User, error) {
	hash, err := hashPassword(userDTO.Password)
	if err != nil {
		return nil, errorsx.New("internal_error", "failed to create user", err)
	}

	entity := mapper.FromUserDTOToUserEntity(userDTO, hash)

	_, err = as.usersRepo.Create(entity)
	if err != nil {
		if errors.Is(err, errorsx.ErrEmailAlreadyExists) {
			return nil, errorsx.New("user_exists", "user with given email already exists", nil)
		}
		return nil, errorsx.New("internal_error", "failed to create user", err)
	}

	return mapper.FromUserEntityToUserDTO(entity), nil
}

func (as *authService) Login(credentials *dto.UserCredentials) (*dto.Token, error) {
	user, err := as.usersRepo.FindByEmail(credentials.Email)
	if err != nil {
		return nil, errorsx.New("internal_error", "failed to find user by email", err)
	}

	if user == nil || !checkPasswordHash(credentials.Password, user.PasswordHash) {
		return nil, errorsx.New("user_not_found", "user with given credentials not found", nil)
	}

	token, err := jwt.GenerateToken(user.Id.String())
	if err != nil {
		return nil, errorsx.New("internal_error", "failed to generate token", err)
	}

	return &dto.Token{Token: token}, nil
}

func (as *authService) GetById(userId uuid.UUID) (*dto.User, error) {
	user, err := as.usersRepo.FindById(userId)
	if err != nil {
		return nil, errorsx.New("internal_error", "failed to find user by id", err)
	}
	if user == nil {
		return nil, errorsx.New("user_not_found", "user with given id not found", nil)
	}

	return mapper.FromUserEntityToUserDTO(user), nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
