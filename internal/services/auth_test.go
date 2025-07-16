package services

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/config"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/stretchr/testify/assert"
)

type MockUsersRepo struct {
	users []*entities.User
}

func (r *MockUsersRepo) Create(user *entities.User) (*entities.User, error) {
	r.users = append(r.users, user)
	return user, nil
}

func (r *MockUsersRepo) FindByEmail(email string) (*entities.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *MockUsersRepo) FindById(userId uuid.UUID) (*entities.User, error) {
	for _, user := range r.users {
		if user.Id == userId {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func TestAuthService_Register(t *testing.T) {
	usersRepo := &MockUsersRepo{}
	service := NewAuthSerice(usersRepo)

	userDTO := &dto.UserDTO{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "11111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := service.Register(userDTO)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(usersRepo.users) != 1 {
		t.Errorf("expected 1 user in repo, but got %d", len(usersRepo.users))
	}
}

func TestAuthService_Login(t *testing.T) {
	config.Load()
	usersRepo := &MockUsersRepo{}
	service := NewAuthSerice(usersRepo)

	userDTO := &dto.UserDTO{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "11111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = service.Register(userDTO)

	_, err := service.Login(&dto.UserCredentials{
		Email:    "test@example.com",
		Password: "11111111"},
	)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	_, err = service.Login(&dto.UserCredentials{
		Email:    "test1@example.com",
		Password: "11111111"},
	)
	if err == nil {
		t.Error("expected error for non-existent user but got none")
	}

	_, err = service.Login(&dto.UserCredentials{
		Email:    "test@example.com",
		Password: "00000000"},
	)
	if err == nil {
		t.Error("expected error for non-existent user but got none")
	}
}

func TestAuthService_GetById(t *testing.T) {
	// Arrange
	usersRepo := &MockUsersRepo{}
	service := NewAuthSerice(usersRepo)

	userDTO := &dto.UserDTO{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "11111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = service.Register(userDTO)

	// Act
	fetchedUser, err := service.GetById(userDTO.Id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, userDTO.Email, fetchedUser.Email)
}
