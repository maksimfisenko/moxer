package repo

import (
	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
)

type TemplatesRepo interface {
	Create(template *entities.Template) (*entities.Template, error)
	FindAllForUser(userID uuid.UUID) ([]*entities.Template, error)
	FindById(id uuid.UUID) (*entities.Template, error)
	FindByNameAndUserId(name string, userId uuid.UUID) (*entities.Template, error)
}
