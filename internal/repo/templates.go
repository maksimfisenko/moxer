package repo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/errorsx"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"gorm.io/gorm"
)

type templatesRepo struct {
	db *gorm.DB
}

func NewTemplatesRepo(db *gorm.DB) *templatesRepo {
	return &templatesRepo{db: db}
}

func (tr *templatesRepo) Create(template *entities.Template) (*entities.Template, error) {
	err := tr.db.Create(template).Error
	if err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return nil, errorsx.ErrInvalidUserId
		}
		return nil, err
	}

	return tr.FindById(template.Id)
}

func (tr *templatesRepo) FindById(id uuid.UUID) (*entities.Template, error) {
	var template entities.Template

	err := tr.db.First(&template, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &template, nil
}

func (tr *templatesRepo) FindAllForUser(userID uuid.UUID) ([]*entities.Template, error) {
	var templates []entities.Template

	err := tr.db.Where("user_id = ?", userID).Find(&templates).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*entities.Template{}, nil
		}
		return nil, err
	}

	templs := make([]*entities.Template, len(templates))
	for i, templ := range templates {
		templs[i] = &templ
	}

	return templs, nil
}
