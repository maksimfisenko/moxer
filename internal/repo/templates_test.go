package repo

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/stretchr/testify/assert"
)

func TestTemplatesRepo_Create(t *testing.T) {
	// Arange
	db, cleanup := setupDB()
	defer cleanup()

	usersRepo := NewUsersRepo(db)
	templatesRepo := NewTemplatesRepo(db)

	user := &entities.User{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "11111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := usersRepo.Create(user)
	if err != nil {
		t.Fatal(err)
	}

	template := &entities.Template{
		Id:   uuid.New(),
		Name: "test_template",
		Content: map[string]any{
			"email":   "{{email}}",
			"country": "{{country}}",
		},
		UserId:    user.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Act
	createdTemplate, err := templatesRepo.Create(template)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, template.Id, createdTemplate.Id)
	assert.Equal(t, template.Name, createdTemplate.Name)
	assert.Equal(t, template.Content, createdTemplate.Content)
	assert.Equal(t, template.UserId, createdTemplate.UserId)
}

func TestTemplatesRepo_findById(t *testing.T) {
	// Arange
	db, cleanup := setupDB()
	defer cleanup()

	usersRepo := NewUsersRepo(db)
	templatesRepo := NewTemplatesRepo(db)

	user := &entities.User{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "11111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = usersRepo.Create(user)

	template := &entities.Template{
		Id:   uuid.New(),
		Name: "test_template",
		Content: map[string]any{
			"email":   "{{email}}",
			"country": "{{country}}",
		},
		UserId:    user.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = templatesRepo.Create(template)

	// Act
	fetchedTemplate, err := templatesRepo.findById(template.Id)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, template.Name, fetchedTemplate.Name)
	assert.Equal(t, template.Content, fetchedTemplate.Content)
	assert.Equal(t, template.UserId, fetchedTemplate.UserId)
}
