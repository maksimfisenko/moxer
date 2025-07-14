package mapper

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/handlers/requests"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/stretchr/testify/assert"
)

func TestFromRegisterRequestToUserDTO(t *testing.T) {
	// Arrange
	req := &requests.RegisterRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	// Act
	dto := FromRegisterRequestToUserDTO(req)

	// Assert
	assert.NotNil(t, dto)
	assert.Equal(t, req.Email, dto.Email)
	assert.NotEqual(t, uuid.Nil, dto.Id)
	assert.WithinDuration(t, time.Now(), dto.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), dto.UpdatedAt, time.Second)
}

func TestFromUserDTOToUserResponse(t *testing.T) {
	// Arrange
	dto := &dto.UserDTO{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Act
	resp := FromUserDTOToUserResponse(dto)

	// Assert
	assert.NotNil(t, resp)
	assert.Equal(t, dto.Id.String(), resp.Id)
	assert.Equal(t, dto.Email, resp.Email)
	assert.Equal(t, dto.CreatedAt, resp.CreatedAt)
	assert.Equal(t, dto.UpdatedAt, resp.UpdatedAt)
}

func TestFromLoginRequestToUserCredentialsDTO(t *testing.T) {
	// Arrange
	req := &requests.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	// Act
	dto := FromLoginRequestToUserCredentialsDTO(req)

	// Assert
	assert.NotNil(t, dto)
	assert.Equal(t, req.Email, dto.Email)
	assert.Equal(t, req.Password, dto.Password)
}

func TestFromTokenDTOToTokenResponse(t *testing.T) {
	// Arrange
	dto := &dto.Token{
		Token: "token",
	}

	// Act
	resp := FromTokenDTOToTokenResponse(dto)

	// Assert
	assert.NotNil(t, resp)
	assert.Equal(t, dto.Token, resp.Token)
}
