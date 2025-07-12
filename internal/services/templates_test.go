package services

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/stretchr/testify/assert"
)

type MockTemplatesRepo struct {
	templates []*entities.Template
}

func (r *MockTemplatesRepo) Create(template *entities.Template) (*entities.Template, error) {
	r.templates = append(r.templates, template)
	return template, nil
}

func (r *MockTemplatesRepo) FindAllForUser(userID uuid.UUID) ([]*entities.Template, error) {
	templs := []*entities.Template{}
	for _, templ := range r.templates {
		if templ.UserId == userID {
			templs = append(templs, templ)
		}
	}
	return templs, nil
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

func TestTemplatesService_GetAllForUser(t *testing.T) {
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

	template1DTO := &dto.Template{
		Id:   uuid.New(),
		Name: "test_template_1",
		Content: map[string]any{
			"email":   "{{email}} 1",
			"country": "{{country}} 1",
		},
		UserId:    userDTO.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	template2DTO := &dto.Template{
		Id:   uuid.New(),
		Name: "test_template_2",
		Content: map[string]any{
			"email":   "{{email}} 2",
			"country": "{{country}} 2",
		},
		UserId:    userDTO.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = authService.Register(userDTO)
	_, _ = templatesService.Create(template1DTO)
	_, _ = templatesService.Create(template2DTO)

	// Act
	fetchedTemplates, err := templatesService.GetAllForUser(userDTO.Id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	if len(fetchedTemplates) != 2 {
		t.Errorf("expected to fetch 2 templates, but got %d", len(fetchedTemplates))
	}

	assert.Equal(t, template1DTO.Name, fetchedTemplates[0].Name)
	assert.Equal(t, template2DTO.Content, fetchedTemplates[1].Content)
}
