package services

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

type MockUsersRepo struct {
	users []*entities.User
}

func (r *MockUsersRepo) Create(user *entities.User) (*entities.User, error) {
	r.users = append(r.users, user)
	return user, nil
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
