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

type MockTemplatesService struct {
	templates map[uuid.UUID]*entities.Template
}

func NewMockTemplatesService() services.TemplatesService {
	return &MockTemplatesService{
		templates: make(map[uuid.UUID]*entities.Template),
	}
}

func (s *MockTemplatesService) Create(templateDTP *dto.Template) (*dto.Template, error) {
	template := mapper.FromTemplateDTOToTemplateEntity(templateDTP)

	s.templates[template.Id] = template

	return mapper.FromTemplateEntityToTemplateDTO(template), nil
}

func TestCreateTemplate(t *testing.T) {
	e := echo.New()
	mockTemplatesService := NewMockTemplatesService()
	handler := handlers.NewTemplatesHandler(e, mockTemplatesService)

	templateReq := requests.CreateTemplateRequest{
		Name: "user",
		Content: map[string]any{
			"email":   "{{email}}",
			"country": "{{country}}",
		},
	}

	templateReqJSON, _ := json.Marshal(templateReq)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/templates", bytes.NewReader(templateReqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", uuid.New().String())

	if err := handler.CreateTemplate(c); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, rec.Code)

	var resp responses.Template
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.NotNil(t, resp.Id)
	assert.Equal(t, templateReq.Name, resp.Name)
	assert.Equal(t, templateReq.Content, resp.Content)
}
