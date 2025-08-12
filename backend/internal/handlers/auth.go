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
//	@Description	Register a new user by given credentials
//	@ID				register
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		requests.CredentialsRequest	true	"New user credentials"
//	@Success		200			{object}	responses.UserResponse		"Sucessfully registered new user"
//	@Failure		400			{object}	errorsx.HTTPError			"Invalid request body"
//	@Failure		409			{object}	errorsx.HTTPError			"User already exists"
//	@Failure		500			{object}	errorsx.HTTPError			"Internal server error"
//	@Router			/public/auth/register [post]
func (ah *authHandler) Register(c echo.Context) error {
	var req requests.CredentialsRequest
	if err := c.Bind(&req); err != nil {
		return errorsx.ErrInvalidRequestBodyHTTP
	}

	dto, err := ah.authService.Register(mapper.FromCredentialsRequestToUserDTO(&req))
	if err != nil {
		switch {
		case errorsx.Is(err, "user_exists"):
			return errorsx.ErrUserExistsHTTP
		default:
			log.Printf("unexpected error: %v", err)
			return errorsx.ErrInternalServerHTTP
		}
	}

	return c.JSON(http.StatusCreated, mapper.FromUserDTOToUserResponse(dto))
}

// Login godoc
//
//	@Summary		Login
//	@Description	Get a JWT token for a user by their credentials
//	@ID				login
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		requests.CredentialsRequest	true	"User's credentials used for logging in"
//	@Success		200			{object}	responses.Token				"Sucessfully logged in a user"
//	@Failure		400			{object}	errorsx.HTTPError			"Invalid request body / User not found"
//	@Failure		500			{object}	errorsx.HTTPError			"Internal server error"
//	@Router			/public/auth/login [post]
func (ah *authHandler) Login(c echo.Context) error {
	var req requests.CredentialsRequest
	if err := c.Bind(&req); err != nil {
		return errorsx.ErrInvalidRequestBodyHTTP
	}

	tokenDTO, err := ah.authService.Login(mapper.FromLoginRequestToUserCredentialsDTO(&req))
	if err != nil {
		switch {
		case errorsx.Is(err, "user_not_found"):
			return errorsx.ErrUserNotFoundHTTP
		default:
			log.Printf("unexpected error: %v", err)
			return errorsx.ErrInternalServerHTTP
		}
	}

	return c.JSON(http.StatusOK, mapper.FromTokenDTOToTokenResponse(tokenDTO))
}

// GetCurrentUser godoc
//
//	@Summary		Get Current User
//	@Description	Get information about current user using a JWT token in Authorization header
//	@ID				me
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.UserResponse	"Sucessfully fetched current user"
//	@Failure		400	{object}	errorsx.HTTPError		"Invalid authentication token / User not found"
//	@Failure		500	{object}	errorsx.HTTPError		"Internal server error"
//	@Router			/private/auth/me [get]
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

	return c.JSON(http.StatusOK, mapper.FromUserDTOToUserResponse(userDTO))
}
