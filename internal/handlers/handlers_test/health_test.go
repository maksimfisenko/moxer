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
	handler := handlers.NewHealthHandler(echo.New())

	req := httptest.NewRequest(http.MethodGet, "/api/v1/healthz", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	if err := handler.HealthCheck(c); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var receivedResp responses.HealthcheckResponse
	err := json.Unmarshal(rec.Body.Bytes(), &receivedResp)
	assert.NoError(t, err)

	assert.Equal(t, "ok", receivedResp.Status)
}
