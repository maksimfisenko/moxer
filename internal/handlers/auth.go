package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maksimfisenko/moxer/internal/handlers/mapper"
	"github.com/maksimfisenko/moxer/internal/handlers/requests"
	"github.com/maksimfisenko/moxer/internal/handlers/responses"
	"github.com/maksimfisenko/moxer/internal/handlers/services"
)

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(e *echo.Echo, authService services.AuthService) *authHandler {
	handler := &authHandler{
		authService: authService,
	}

	e.POST("/api/v1/auth/register", handler.Register)
	e.POST("/api/v1/auth/login", handler.Login)
	e.GET("/api/v1/auth/me", handler.GetCurrentUser)

	return handler
}

// Register godoc
//
//	@Summary		Register
//	@Description	Register new user by given credentials (email, password)
//	@ID				register
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		requests.RegisterRequest	true	"Register request"
//	@Success		200			{object}	responses.UserResponse		"Sucessfully registered new user"
//	@Failure		400			{object}	responses.ErrorResponse		"Failed to parse request body"
//	@Failure		500			{object}	responses.ErrorResponse		"Failed to register"
//	@Router			/auth/register [post]
func (ah *authHandler) Register(c echo.Context) error {
	var req requests.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "failed to parse request body",
		})
	}

	dto := mapper.FromRegisterRequestToUserDTO(&req)

	dto, err := ah.authService.Register(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: fmt.Sprintf("failed to register: %v", err),
		})
	}

	resp := mapper.FromUserDTOToUserResponse(dto)

	return c.JSON(http.StatusCreated, resp)
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login new user by given credentials (email, password)
//	@ID				login
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		requests.LoginRequest	true	"Login request"
//	@Success		200			{object}	responses.Token			"Sucessfully registered new user"
//	@Failure		400			{object}	responses.ErrorResponse	"Failed to parse request body"
//	@Failure		500			{object}	responses.ErrorResponse	"Failed to login"
//	@Router			/auth/login [post]
func (ah *authHandler) Login(c echo.Context) error {
	var req requests.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "failed to parse request body",
		})
	}

	credentialsDTO := mapper.FromLoginRequestToUserCredentialsDTO(&req)

	tokenDTO, err := ah.authService.Login(credentialsDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: fmt.Sprintf("failed to register: %v", err),
		})
	}

	resp := mapper.FromTokenDTOToTokenResponse(tokenDTO)

	return c.JSON(http.StatusOK, resp)
}

// GetCurrentUser godoc
//
//	@Summary		Get current user
//	@Description	Get current user by JWT token in Authorization header
//	@ID				me
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.UserResponse	"Sucessfully fetched current user"
//	@Failure		400	{object}	responses.ErrorResponse	"Failed to parse token"
//	@Failure		500	{object}	responses.ErrorResponse	"Failed to fetch current user"
//	@Router			/auth/me [get]
func (ah *authHandler) GetCurrentUser(c echo.Context) error {
	userIdRaw := c.Get("userId").(string)

	userId, err := uuid.Parse(userIdRaw)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "failed to parse token",
		})
	}

	userDTO, err := ah.authService.GetById(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: fmt.Sprintf("failed to fetch user: %v", err),
		})
	}

	resp := mapper.FromUserDTOToUserResponse(userDTO)

	return c.JSON(http.StatusOK, resp)
}
