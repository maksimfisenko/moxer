package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/errorsx"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/maksimfisenko/moxer/internal/services/generator"
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

	_, err := ts.templatesRepo.Create(entity)
	if err != nil {
		if errors.Is(err, errorsx.ErrInvalidUserId) {
			return nil, errorsx.New("user_not_found", "user with given id not found", nil)
		}
		return nil, errorsx.New("internal_error", "failed to create template", err)
	}

	return mapper.FromTemplateEntityToTemplateDTO(entity), nil
}

func (ts *templatesService) GetAllForUser(userID uuid.UUID) ([]*dto.Template, error) {
	templates, err := ts.templatesRepo.FindAllForUser(userID)
	if err != nil {
		return nil, errorsx.New("internal_error", "failed to fetch templates for user", err)
	}

	return mapper.FromTemplateEntityListToTemplateDTOList(templates), nil
}

func (ts *templatesService) GenerateData(templateId uuid.UUID, count int) (*dto.GeneratedData, error) {
	template, err := ts.templatesRepo.FindById(templateId)
	if err != nil {
		return nil, errorsx.New("internal_error", "failed to find template by id", err)
	}
	if template == nil {
		return nil, errorsx.New("template_not_found", "template with given id not found", nil)
	}

	return &dto.GeneratedData{
		Data: generator.GenerateData(template.Content, count),
	}, nil
}
