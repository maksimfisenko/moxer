package repo

import (
	"github.com/google/uuid"
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
	if err := tr.db.Create(template).Error; err != nil {
		return nil, err
	}

	return tr.findById(template.Id)
}

func (tr *templatesRepo) findById(id uuid.UUID) (*entities.Template, error) {
	var template entities.Template
	if err := tr.db.First(&template, id).Error; err != nil {
		return nil, err
	}
	return &template, nil
}
