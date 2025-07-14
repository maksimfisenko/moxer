package mapper

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/stretchr/testify/assert"
)

func TestFromTemplateDTOToTemplateEntity(t *testing.T) {
	// Arrange
	dto := &dto.Template{
		Id:     uuid.New(),
		UserId: uuid.New(),
		Name:   "name",
		Content: map[string]any{
			"field_1": "value_1",
			"fields_2_3": map[string]string{
				"field_2": "value_2",
				"field_3": "value_3",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Act
	entity := FromTemplateDTOToTemplateEntity(dto)

	// Assert
	assert.NotNil(t, entity)
	assert.Equal(t, dto.Id, entity.Id)
	assert.Equal(t, dto.Name, entity.Name)
	assert.Equal(t, dto.Content, entity.Content)
	assert.Equal(t, dto.CreatedAt, entity.CreatedAt)
	assert.Equal(t, dto.UpdatedAt, entity.UpdatedAt)
}

func TestFromTemplateEntityToTemplateDTO(t *testing.T) {
	// Arrange
	entity := &entities.Template{
		Id:     uuid.New(),
		UserId: uuid.New(),
		Name:   "name",
		Content: map[string]any{
			"field_1": "value_1",
			"fields_2_3": map[string]string{
				"field_2": "value_2",
				"field_3": "value_3",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Act
	dto := FromTemplateEntityToTemplateDTO(entity)

	// Assert
	assert.NotNil(t, dto)
	assert.Equal(t, entity.Id, dto.Id)
	assert.Equal(t, entity.Name, dto.Name)
	assert.Equal(t, entity.Content, dto.Content)
	assert.Equal(t, entity.CreatedAt, dto.CreatedAt)
	assert.Equal(t, entity.UpdatedAt, dto.UpdatedAt)
}

func TestFromTemplateEntityListToTemplateDTOList(t *testing.T) {
	// Arrange
	entity1 := &entities.Template{
		Id:     uuid.New(),
		UserId: uuid.New(),
		Name:   "name_1",
		Content: map[string]any{
			"field_1": "value_1",
			"fields_2_3": map[string]string{
				"field_2": "value_2",
				"field_3": "value_3",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	entity2 := &entities.Template{
		Id:     uuid.New(),
		UserId: uuid.New(),
		Name:   "name_2",
		Content: map[string]any{
			"field_4": "value_4",
			"fields_5_6": map[string]string{
				"field_5": "value_5",
				"field_6": "value_6",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	entitiesList := []*entities.Template{entity1, entity2}

	// Act
	dtoList := FromTemplateEntityListToTemplateDTOList(entitiesList)

	// Assert
	assert.NotNil(t, dtoList)
	assert.Equal(t, len(entitiesList), len(dtoList))
	assert.Equal(t, entity1.Id, dtoList[0].Id)
	assert.Equal(t, entity2.Name, dtoList[1].Name)
	assert.Equal(t, entity1.Content, dtoList[0].Content)
	assert.Equal(t, entity2.CreatedAt, dtoList[1].CreatedAt)
	assert.Equal(t, entity1.UpdatedAt, dtoList[0].UpdatedAt)
}
