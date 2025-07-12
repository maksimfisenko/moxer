package services

import (
	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

type TemplatesService interface {
	Create(templateDTO *dto.Template) (*dto.Template, error)
	GetAllForUser(userID uuid.UUID) ([]*dto.Template, error)
}
