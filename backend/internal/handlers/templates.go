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

func NewTemplatesHandler(private *echo.Group, templatesService services.TemplatesService) *templatesHandler {
	handler := &templatesHandler{
		templatesService: templatesService,
	}

	private.POST("/api/v1/templates", handler.CreateTemplate)
	private.GET("/api/v1/templates", handler.GetAllForUser)
	private.POST("/api/v1/templates/:id/generate", handler.GenerateData)

	return handler
}

// CreateTemplate godoc
//
//	@Summary		Create Template
//	@Description	Create a new template with provided request body and user id
//	@ID				create-template
//	@Tags			templates
//	@Accept			json
//	@Produce		json
//	@Param			data	body		requests.CreateTemplateRequest	true	"Create template request"
//	@Success		200		{object}	responses.Template				"Sucessfully created new template"
//	@Failure		400		{object}	errorsx.HTTPError				"Invalid authentication token / Invalid request body / User not found"
//	@Failure		409		{object}	errorsx.HTTPError				"Template with given name already exists"
//	@Failure		500		{object}	errorsx.HTTPError				"Internal server error"
//	@Router			/private/templates [post]
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

	dto, err := th.templatesService.Create(mapper.FromCreateTemplateRequestToTemplateDTO(&req, userId))
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

	return c.JSON(http.StatusCreated, mapper.FromTemplateDTOToTemplateResponse(dto))
}

// GetAllForUser godoc
//
//	@Summary		Get User's Template
//	@Description	Get all templates for a certain user by their JWT token
//	@ID				get-all-for-user
//	@Tags			templates
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		responses.Template	"Successfully fetched user's templates"
//	@Failure		400	{object}	errorsx.HTTPError	"Invalid authentication token"
//	@Failure		500	{object}	errorsx.HTTPError	"Internal server error"
//	@Router			/private/templates [get]
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

	return c.JSON(http.StatusOK, mapper.FromTemplateDTOListToTemplateResponseList(dtoList))
}

// GenerateData godoc
//
//	@Summary		Generate Data
//	@Description	Generate data for a selected template
//	@ID				generate-data
//	@Tags			templates
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		responses.GeneratedData	"Successfully generated data"
//	@Failure		400	{object}	errorsx.HTTPError		"Invalid template id / Invalid request body / Template not found"
//	@Failure		500	{object}	errorsx.HTTPError		"Internal server error"
//	@Router			/private/templates/:id/generate [post]
func (th *templatesHandler) GenerateData(c echo.Context) error {
	templateId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return errorsx.ErrInvalidTemplateIdHTTP
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

	return c.JSON(http.StatusCreated, mapper.FromGeneratedDataDTOToGeneratedDataResponse(dto))
}
