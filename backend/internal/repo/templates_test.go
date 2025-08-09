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
		Id:           uuid.New(),
		Email:        "test@example.com",
		PasswordHash: "password_hash",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
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
		Id:           uuid.New(),
		Email:        "test@example.com",
		PasswordHash: "password_hash",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
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
	fetchedTemplate, err := templatesRepo.FindById(template.Id)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, template.Name, fetchedTemplate.Name)
	assert.Equal(t, template.Content, fetchedTemplate.Content)
	assert.Equal(t, template.UserId, fetchedTemplate.UserId)
}

func TestTemplatesRepo_FindALlForUser(t *testing.T) {
	// Arange
	db, cleanup := setupDB()
	defer cleanup()

	usersRepo := NewUsersRepo(db)
	templatesRepo := NewTemplatesRepo(db)

	user := &entities.User{
		Id:           uuid.New(),
		Email:        "test@example.com",
		PasswordHash: "password_hash",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, _ = usersRepo.Create(user)

	template1 := &entities.Template{
		Id:   uuid.New(),
		Name: "test_template_1",
		Content: map[string]any{
			"email":   "{{email}} 1",
			"country": "{{country}} 1",
		},
		UserId:    user.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	template2 := &entities.Template{
		Id:   uuid.New(),
		Name: "test_template_2",
		Content: map[string]any{
			"email":   "{{email}} 2",
			"country": "{{country}} 2",
		},
		UserId:    user.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = templatesRepo.Create(template1)
	_, _ = templatesRepo.Create(template2)

	// Act
	fetchedTemplates, err := templatesRepo.FindAllForUser(user.Id)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 2, len(fetchedTemplates))
	assert.Equal(t, template1.Name, fetchedTemplates[0].Name)
	assert.Equal(t, template2.Content, fetchedTemplates[1].Content)
}

func TestTemplatesRepo_FindByNameAndUserId(t *testing.T) {
	// Arange
	db, cleanup := setupDB()
	defer cleanup()

	usersRepo := NewUsersRepo(db)
	templatesRepo := NewTemplatesRepo(db)

	user := &entities.User{
		Id:           uuid.New(),
		Email:        "test@example.com",
		PasswordHash: "password_hash",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
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
	fetchedTemplate, err := templatesRepo.FindByNameAndUserId(template.Name, template.UserId)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, template.Id, fetchedTemplate.Id)
	assert.Equal(t, template.Content, fetchedTemplate.Content)
}
