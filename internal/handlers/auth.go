package handlers

import (
	"fmt"
	"net/http"

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
