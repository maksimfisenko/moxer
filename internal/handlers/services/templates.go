package services

import (
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

type TemplatesService interface {
	Create(templateDTO *dto.Template) (*dto.Template, error)
}
