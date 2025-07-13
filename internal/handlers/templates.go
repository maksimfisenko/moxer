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

type templatesHandler struct {
	templatesService services.TemplatesService
}

func NewTemplatesHandler(e *echo.Echo, templatesService services.TemplatesService) *templatesHandler {
	handler := &templatesHandler{
		templatesService: templatesService,
	}

	e.POST("/api/v1/templates", handler.CreateTemplate)
	e.GET("/api/v1/templates", handler.GetAllForUser)
	e.POST("/api/v1/templates/:id/generate", handler.GenerateData)

	return handler
}

// CreateTemplate godoc
//
//	@Summary		Create template
//	@Description	Create a new template with provided request body and user id
//	@ID				create-template
//	@Tags			templates
//	@Accept			json
//	@Produce		json
//	@Param			data	body		requests.CreateTemplateRequest	true	"Create template request"
//	@Success		200		{object}	responses.Template				"Sucessfully created new template"
//	@Failure		400		{object}	responses.ErrorResponse			"Failed to parse request body or token"
//	@Failure		500		{object}	responses.ErrorResponse			"Failed to create new template"
//	@Router			/templates [post]
func (th *templatesHandler) CreateTemplate(c echo.Context) error {
	userIdRaw := c.Get("userId").(string)

	userId, err := uuid.Parse(userIdRaw)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "failed to parse token",
		})
	}

	var req requests.CreateTemplateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "failed to parse request body",
		})
	}

	dto := mapper.FromCreateTemplateRequestToTemplateDTO(&req, userId)

	dto, err = th.templatesService.Create(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: fmt.Sprintf("failed to create new template: %v", err),
		})
	}

	resp := mapper.FromTemplateDTOToTemplateResponse(dto)

	return c.JSON(http.StatusCreated, resp)
}

// GetAllForUser godoc
//
//	@Summary		Get user's templates
//	@Description	Get all templates of certain user by given JWT token
//	@ID				get-all-for-user
//	@Tags			templates
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		responses.Template		"Successfully fetched user's templates"
//	@Failure		400	{object}	responses.ErrorResponse	"Failed to parse token"
//	@Failure		500	{object}	responses.ErrorResponse	"Failed to fetch user's templates"
//	@Router			/templates [get]
func (th *templatesHandler) GetAllForUser(c echo.Context) error {
	userIdRaw := c.Get("userId").(string)

	userId, err := uuid.Parse(userIdRaw)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "failed to parse token",
		})
	}

	dtoList, err := th.templatesService.GetAllForUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: fmt.Sprintf("failed to fetch user's templates: %v", err),
		})
	}

	resp := mapper.FromTemplateDTOListToTemplateResponseList(dtoList)

	return c.JSON(http.StatusOK, resp)
}

// GenerateData godoc
//
//	@Summary		Generate data
//	@Description	Generate data from template
//	@ID				generate-data
//	@Tags			templates
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		responses.Template		"Successfully generated data from template"
//	@Failure		400	{object}	responses.ErrorResponse	"Failed to parse param / request"
//	@Failure		500	{object}	responses.ErrorResponse	"Failed to generate data"
//	@Router			/templates/:id/generate [post]
func (th *templatesHandler) GenerateData(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "failed to parse template id param",
		})
	}

	var req requests.GenerateDataRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Error: "failed to parse request body",
		})
	}

	dto, err := th.templatesService.GenerateData(id, req.Count)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: fmt.Sprintf("failed to generate data: %v", err),
		})
	}

	resp := mapper.FromGeneratedDataDTOToGeneratedDataResponse(dto)

	return c.JSON(http.StatusCreated, resp)
}
