package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maksimfisenko/moxer/internal/handlers"
	mapper2 "github.com/maksimfisenko/moxer/internal/handlers/mapper"
	"github.com/maksimfisenko/moxer/internal/handlers/requests"
	"github.com/maksimfisenko/moxer/internal/handlers/responses"
	"github.com/maksimfisenko/moxer/internal/handlers/services"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/maksimfisenko/moxer/internal/services/mapper"
	"github.com/stretchr/testify/assert"
)

type MockAuthService struct {
	users map[uuid.UUID]*entities.User
}

func NewMockAuthService() services.AuthService {
	return &MockAuthService{
		users: make(map[uuid.UUID]*entities.User),
	}
}

func (s *MockAuthService) Register(userDTO *dto.UserDTO) (*dto.UserDTO, error) {
	user := mapper.FromUserDTOToUserEntity(userDTO)

	s.users[user.Id] = user

	return mapper.FromUserEntityToUserDTO(user), nil
}

func (s *MockAuthService) Login(credentials *dto.UserCredentials) (*dto.Token, error) {
	for _, user := range s.users {
		if user.Email == credentials.Email && user.Password == credentials.Password {
			return &dto.Token{Token: "12345"}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *MockAuthService) GetById(userId uuid.UUID) (*dto.UserDTO, error) {
	user, ok := s.users[userId]
	if !ok {
		return nil, errors.New("user not found")
	}
	return mapper.FromUserEntityToUserDTO(user), nil
}

func TestRegister(t *testing.T) {
	mockAuthService := NewMockAuthService()
	handler := handlers.NewAuthHandler(echo.New(), mockAuthService)

	registerReq := requests.RegisterRequest{
		Email:    "email@example.com",
		Password: "11111111",
	}

	registerReqJSON, _ := json.Marshal(registerReq)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewReader(registerReqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	if err := handler.Register(c); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, rec.Code)

	var resp responses.UserResponse
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)

	assert.NotNil(t, resp.Id)
	assert.Equal(t, registerReq.Email, resp.Email)
}

func TestLogin(t *testing.T) {
	// Arrange
	mockAuthService := NewMockAuthService()
	handler := handlers.NewAuthHandler(echo.New(), mockAuthService)

	registerReq := requests.RegisterRequest{
		Email:    "email@example.com",
		Password: "11111111",
	}
	userDTO := mapper2.FromRegisterRequestToUserDTO(&registerReq)

	loginReq := requests.LoginRequest{
		Email:    "email@example.com",
		Password: "11111111",
	}

	_, err := mockAuthService.Register(userDTO)
	assert.NoError(t, err)

	loginReqJSON, _ := json.Marshal(loginReq)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(loginReqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	// Act
	if err := handler.Login(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp responses.Token
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)

	assert.Equal(t, "12345", resp.Token)
}

func TestGetCurrentUser(t *testing.T) {
	// Arrange
	mockAuthService := NewMockAuthService()
	handler := handlers.NewAuthHandler(echo.New(), mockAuthService)

	registerReq := requests.RegisterRequest{
		Email:    "email@example.com",
		Password: "11111111",
	}
	userDTO := mapper2.FromRegisterRequestToUserDTO(&registerReq)

	_, err := mockAuthService.Register(userDTO)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/auth/me", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.Set("userId", userDTO.Id.String())

	// Act
	if err := handler.GetCurrentUser(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp responses.UserResponse
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)

	assert.NotNil(t, resp)
	assert.Equal(t, userDTO.Email, resp.Email)
}
