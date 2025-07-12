package mapper

import (
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

func FromTemplateDTOToTemplateEntity(dto *dto.Template) *entities.Template {
	return &entities.Template{
		Id:        dto.Id,
		Name:      dto.Name,
		Content:   dto.Content,
		UserId:    dto.UserId,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

func FromTemplateEntityToTemplateDTO(entity *entities.Template) *dto.Template {
	return &dto.Template{
		Id:        entity.Id,
		Name:      entity.Name,
		Content:   entity.Content,
		UserId:    entity.UserId,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
