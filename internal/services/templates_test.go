package services

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

type MockTemplatesRepo struct {
	templates []*entities.Template
}

func (r *MockTemplatesRepo) Create(template *entities.Template) (*entities.Template, error) {
	r.templates = append(r.templates, template)
	return template, nil
}

func TestTemplatesService_Create(t *testing.T) {
	// Arrange
	usersRepo := &MockUsersRepo{}
	templatesRepo := &MockTemplatesRepo{}

	authService := NewAuthSerice(usersRepo)
	templatesService := NewTemplatesService(templatesRepo)

	userDTO := &dto.UserDTO{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "11111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	templateDTP := &dto.Template{
		Id:   uuid.New(),
		Name: "test_template",
		Content: map[string]any{
			"email":   "{{email}}",
			"country": "{{country}}",
		},
		UserId:    userDTO.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = authService.Register(userDTO)

	// Act
	_, err := templatesService.Create(templateDTP)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(templatesRepo.templates) != 1 {
		t.Errorf("expected 1 template in repo, but got %d", len(templatesRepo.templates))
	}
}
