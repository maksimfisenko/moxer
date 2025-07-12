package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/maksimfisenko/moxer/internal/handlers"
	"github.com/maksimfisenko/moxer/internal/handlers/responses"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Arrange
	e := echo.New()
	handler := handlers.NewHealthHandler(e)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/healthz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act
	if err := handler.HealthCheck(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp responses.HealthcheckResponse
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "ok", resp.Status)
}
