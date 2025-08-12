package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maksimfisenko/moxer/internal/errorsx"
	"github.com/maksimfisenko/moxer/internal/handlers/mapper"
	"github.com/maksimfisenko/moxer/internal/handlers/requests"
	"github.com/maksimfisenko/moxer/internal/handlers/services"
)

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(public, private *echo.Group, authService services.AuthService) *authHandler {
	handler := &authHandler{
		authService: authService,
	}

	public.POST("/auth/register", handler.Register)
	public.POST("/auth/login", handler.Login)

	private.GET("/auth/me", handler.GetCurrentUser)

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
		return errorsx.ErrInvalidRequestBodyHTTP
	}

	dto := mapper.FromRegisterRequestToUserDTO(&req)

	dto, err := ah.authService.Register(dto)
	if err != nil {
		switch {
		case errorsx.Is(err, "user_exists"):
			return errorsx.ErrUserExistsHTTP
		default:
			log.Printf("unexpected error: %v", err)
			return errorsx.ErrInternalServerHTTP
		}
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
		return errorsx.ErrInvalidRequestBodyHTTP
	}

	credentialsDTO := mapper.FromLoginRequestToUserCredentialsDTO(&req)

	tokenDTO, err := ah.authService.Login(credentialsDTO)
	if err != nil {
		switch {
		case errorsx.Is(err, "user_not_found"):
			return errorsx.ErrUserNotFoundHTTP
		default:
			log.Printf("unexpected error: %v", err)
			return errorsx.ErrInternalServerHTTP
		}
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
		return errorsx.ErrInvalidTokenHTTP
	}

	userDTO, err := ah.authService.GetById(userId)
	if err != nil {
		switch {
		case errorsx.Is(err, "user_not_found"):
			return errorsx.ErrUserNotFoundHTTP
		default:
			log.Printf("unexpected error: %v", err)
			return errorsx.ErrInternalServerHTTP
		}
	}

	resp := mapper.FromUserDTOToUserResponse(userDTO)

	return c.JSON(http.StatusOK, resp)
}
