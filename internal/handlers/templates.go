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
		return errorsx.ErrInvalidTokenHTTP
	}

	var req requests.CreateTemplateRequest
	if err := c.Bind(&req); err != nil {
		return errorsx.ErrInvalidRequestBodyHTTP
	}

	dto := mapper.FromCreateTemplateRequestToTemplateDTO(&req, userId)

	dto, err = th.templatesService.Create(dto)
	if err != nil {
		switch {
		case errorsx.Is(err, "template_exists"):
			return errorsx.ErrTemplateExistsHTTP
		case errorsx.Is(err, "user_not_found"):
			return errorsx.ErrUserNotFoundHTTP
		default:
			log.Printf("unexpected error: %v", err)
			return errorsx.ErrInternalServerHTTP
		}
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
		return errorsx.ErrInvalidTokenHTTP
	}

	dtoList, err := th.templatesService.GetAllForUser(userId)
	if err != nil {
		log.Printf("unexpected error: %v", err)
		return errorsx.ErrInternalServerHTTP
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
	templateId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errorsx.ErrInvalidTokenHTTP
	}

	var req requests.GenerateDataRequest
	if err := c.Bind(&req); err != nil {
		return errorsx.ErrInvalidRequestBodyHTTP
	}

	dto, err := th.templatesService.GenerateData(templateId, req.Count)
	if err != nil {
		switch {
		case errorsx.Is(err, "template_not_found"):
			return errorsx.ErrTemplateNotFoundHTTP
		default:
			log.Printf("unexpected error: %v", err)
			return errorsx.ErrInternalServerHTTP
		}
	}

	resp := mapper.FromGeneratedDataDTOToGeneratedDataResponse(dto)

	return c.JSON(http.StatusCreated, resp)
}
