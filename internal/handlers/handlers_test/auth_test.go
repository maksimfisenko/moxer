package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maksimfisenko/moxer/internal/handlers"
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
