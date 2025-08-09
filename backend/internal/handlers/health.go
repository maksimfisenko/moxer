package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maksimfisenko/moxer/internal/handlers/responses"
)

type healthHandler struct{}

func NewHealthHandler(e *echo.Echo) *healthHandler {
	handler := &healthHandler{}

	e.GET("/api/v1/healthz", handler.HealthCheck)

	return handler
}

// HealthCheck godoc
//
//	@Summary		Health Check
//	@Description	Check if the application is up
//	@ID				health-check
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	responses.HealthcheckResponse "Sucessfully received response from server"
//	@Router			/healthz [get]
func (hh *healthHandler) HealthCheck(c echo.Context) error {
	resp := responses.HealthcheckResponse{Status: "ok"}
	return c.JSON(http.StatusOK, resp)
}
