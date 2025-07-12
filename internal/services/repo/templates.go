package repo

import (
	"github.com/maksimfisenko/moxer/internal/repo/entities"
)

type TemplatesRepo interface {
	Create(template *entities.Template) (*entities.Template, error)
}
