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
//	@Success		200			{object}	responses.Template		"Sucessfully created new template"
//	@Failure		400			{object}	responses.ErrorResponse		"Failed to parse request body or token"
//	@Failure		500			{object}	responses.ErrorResponse		"Failed to create new template"
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
