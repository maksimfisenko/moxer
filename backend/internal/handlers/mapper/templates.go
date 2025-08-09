package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/handlers/requests"
	"github.com/maksimfisenko/moxer/internal/handlers/responses"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

func FromCreateTemplateRequestToTemplateDTO(req *requests.CreateTemplateRequest, userId uuid.UUID) *dto.Template {
	return &dto.Template{
		Id:        uuid.New(),
		UserId:    userId,
		Name:      req.Name,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func FromTemplateDTOToTemplateResponse(dto *dto.Template) *responses.Template {
	return &responses.Template{
		Id:        dto.Id.String(),
		Name:      dto.Name,
		Content:   dto.Content,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

func FromTemplateDTOListToTemplateResponseList(dtoList []*dto.Template) []*responses.Template {
	responsses := make([]*responses.Template, len(dtoList))
	for i, dto := range dtoList {
		responsses[i] = FromTemplateDTOToTemplateResponse(dto)
	}
	return responsses
}

func FromGeneratedDataDTOToGeneratedDataResponse(dto *dto.GeneratedData) *responses.GeneratedData {
	return &responses.GeneratedData{
		Data: dto.Data,
	}
}
