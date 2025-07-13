package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	"github.com/maksimfisenko/moxer/internal/services/generator"
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

func (s *MockTemplatesService) GetAllForUser(userId uuid.UUID) ([]*dto.Template, error) {
	templates := []*entities.Template{}
	for _, template := range s.templates {
		if template.UserId == userId {
			templates = append(templates, template)
		}
	}
	return mapper.FromTemplateEntityListToTemplateDTOList(templates), nil
}

func (s *MockTemplatesService) GenerateData(templateId uuid.UUID, count int) (*dto.GeneratedData, error) {
	templ, ok := s.templates[templateId]
	if !ok {
		return nil, errors.New("template not found")
	}

	return &dto.GeneratedData{
		Data: generator.GenerateData(templ.Content, count),
	}, nil
}

func TestCreateTemplate(t *testing.T) {
	// Arrange
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

	// Act
	if err := handler.CreateTemplate(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusCreated, rec.Code)

	var resp responses.Template
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.NotNil(t, resp.Id)
	assert.Equal(t, templateReq.Name, resp.Name)
	assert.Equal(t, templateReq.Content, resp.Content)
}

func TestGetAllForUser(t *testing.T) {
	// Arrange
	e := echo.New()
	mockTemplatesService := NewMockTemplatesService()
	handler := handlers.NewTemplatesHandler(e, mockTemplatesService)

	template1Req := requests.CreateTemplateRequest{
		Name: "user_1",
		Content: map[string]any{
			"email":   "{{email}}",
			"country": "{{country}}",
		},
	}

	template2Req := requests.CreateTemplateRequest{
		Name: "user_2",
		Content: map[string]any{
			"email":   "{{email}}",
			"country": "{{country}}",
		},
	}

	template1ReqJSON, _ := json.Marshal(template1Req)
	template2ReqJSON, _ := json.Marshal(template2Req)

	userId := uuid.New().String()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/templates", bytes.NewReader(template1ReqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", userId)

	if err := handler.CreateTemplate(c); err != nil {
		t.Fatal(err)
	}

	req = httptest.NewRequest(http.MethodPost, "/api/v1/templates", bytes.NewReader(template2ReqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.Set("userId", userId)

	if err := handler.CreateTemplate(c); err != nil {
		t.Fatal(err)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/v1/templates", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.Set("userId", userId)

	// Act
	if err := handler.GetAllForUser(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp []responses.Template
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.NotNil(t, resp[0].Id)
	assert.NotNil(t, resp[1].Id)
	assert.Equal(t, template1Req.Name, resp[0].Name)
	assert.Equal(t, template2Req.Content, resp[1].Content)
}

func TestGenerateData(t *testing.T) {
	// Arrange
	e := echo.New()
	mockTemplatesService := NewMockTemplatesService()
	handler := handlers.NewTemplatesHandler(e, mockTemplatesService)

	templateReq := requests.CreateTemplateRequest{
		Name: "user",
		Content: map[string]any{
			"uuid": "{{email}}",
			"name": "{{name}}",
		},
	}

	templateReqJSON, _ := json.Marshal(templateReq)
	userId := uuid.New().String()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/templates", bytes.NewReader(templateReqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("userId", userId)

	if err := handler.CreateTemplate(c); err != nil {
		t.Fatal(err)
	}

	var templ responses.Template
	err := json.Unmarshal(rec.Body.Bytes(), &templ)
	assert.NoError(t, err)

	url := fmt.Sprintf("/api/v1/templates/%s/generate", templ.Id)

	fmt.Println(url)

	genReq := requests.GenerateDataRequest{
		Count: 3,
	}
	genReqJSON, _ := json.Marshal(genReq)

	req = httptest.NewRequest(http.MethodPost, url, bytes.NewReader(genReqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.Set("userId", userId)

	c.SetPath("/api/v1/templates/:id/generate")
	c.SetParamNames("id")
	c.SetParamValues(templ.Id)

	// Act
	if err := handler.GenerateData(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusCreated, rec.Code)

	var resp responses.GeneratedData
	err = json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.NotNil(t, resp)
	assert.Equal(t, genReq.Count, len(resp.Data))
}
