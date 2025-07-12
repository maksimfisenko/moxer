package services

import (
	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/maksimfisenko/moxer/internal/services/mapper"
	"github.com/maksimfisenko/moxer/internal/services/repo"
)

type templatesService struct {
	templatesRepo repo.TemplatesRepo
}

func NewTemplatesService(templatesRepo repo.TemplatesRepo) *templatesService {
	return &templatesService{templatesRepo: templatesRepo}
}

func (ts *templatesService) Create(templateDTO *dto.Template) (*dto.Template, error) {
	entity := mapper.FromTemplateDTOToTemplateEntity(templateDTO)

	if _, err := ts.templatesRepo.Create(entity); err != nil {
		return nil, err
	}

	return mapper.FromTemplateEntityToTemplateDTO(entity), nil
}

func (ts *templatesService) GetAllForUser(userID uuid.UUID) ([]*dto.Template, error) {
	templates, err := ts.templatesRepo.FindAllForUser(userID)
	if err != nil {
		return nil, err
	}

	return mapper.FromTemplateEntityListToTemplateDTOList(templates), nil
}
