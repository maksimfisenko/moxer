package mapper

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/stretchr/testify/assert"
)

func TestUserDTOToUserEntity(t *testing.T) {
	// Arrange
	dto := &dto.User{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hash := "password_hash"

	// Act
	entity := FromUserDTOToUserEntity(dto, hash)

	// Assert
	assert.NotNil(t, entity)
	assert.Equal(t, dto.Id, entity.Id)
	assert.Equal(t, dto.Email, entity.Email)
	assert.Equal(t, dto.CreatedAt, entity.CreatedAt)
	assert.Equal(t, dto.UpdatedAt, entity.UpdatedAt)
}

func TestUserEntityToUserDTO(t *testing.T) {
	// Arrange
	entity := &entities.User{
		Id:           uuid.New(),
		Email:        "test@example.com",
		PasswordHash: "password_hash",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Act
	dto := FromUserEntityToUserDTO(entity)

	// Assert
	assert.NotNil(t, dto)
	assert.Equal(t, entity.Id, dto.Id)
	assert.Equal(t, entity.Email, dto.Email)
	assert.Equal(t, entity.CreatedAt, dto.CreatedAt)
	assert.Equal(t, entity.UpdatedAt, dto.UpdatedAt)
}
