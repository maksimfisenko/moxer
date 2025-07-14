package mapper

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/handlers/requests"
	"github.com/maksimfisenko/moxer/internal/services/dto"
	"github.com/stretchr/testify/assert"
)

func TestFromCreateTemplateRequestToTemplateDTO(t *testing.T) {
	// Arrange
	req := &requests.CreateTemplateRequest{
		Name: "name",
		Content: map[string]any{
			"field_1": "value_1",
			"fields_2_3": map[string]string{
				"field_2": "value_2",
				"field_3": "value_3",
			},
		},
	}

	userID := uuid.New()

	// Act
	dto := FromCreateTemplateRequestToTemplateDTO(req, userID)

	// Assert
	assert.NotNil(t, dto)
	assert.Equal(t, req.Name, dto.Name)
	assert.Equal(t, req.Content, dto.Content)
	assert.Equal(t, userID, dto.UserId)
	assert.NotEqual(t, uuid.Nil, dto.Id)
	assert.WithinDuration(t, time.Now(), dto.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), dto.UpdatedAt, time.Second)
}

func TestFromTemplateDTOToTemplateResponse(t *testing.T) {
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
	resp := FromTemplateDTOToTemplateResponse(dto)

	// Assert
	assert.NotNil(t, resp)
	assert.Equal(t, dto.Id.String(), resp.Id)
	assert.Equal(t, dto.Name, resp.Name)
	assert.Equal(t, dto.Content, resp.Content)
	assert.Equal(t, dto.CreatedAt, resp.CreatedAt)
	assert.Equal(t, dto.UpdatedAt, resp.UpdatedAt)
}

func TestFromTemplateDTOListToTemplateResponseList(t *testing.T) {
	// Arrange
	dto1 := &dto.Template{
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

	dto2 := &dto.Template{
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

	dtoList := []*dto.Template{dto1, dto2}

	// Act
	resp := FromTemplateDTOListToTemplateResponseList(dtoList)

	// Assert
	assert.NotNil(t, resp)
	assert.Equal(t, len(dtoList), len(resp))
	assert.Equal(t, dto1.Id.String(), resp[0].Id)
	assert.Equal(t, dto2.Name, resp[1].Name)
	assert.Equal(t, dto1.Content, resp[0].Content)
	assert.Equal(t, dto2.CreatedAt, resp[1].CreatedAt)
	assert.Equal(t, dto1.UpdatedAt, resp[0].UpdatedAt)
}

func TestFromGeneratedDataDTOToGeneratedDataResponse(t *testing.T) {
	// Arrange
	obj1 := map[string]any{
		"field_1": "value_1",
		"fields_2_3": map[string]string{
			"field_2": "value_2",
			"field_3": "value_3",
		},
	}

	obj2 := map[string]any{
		"field_4": "value_4",
		"fields_5_6": map[string]string{
			"field_5": "value_5",
			"field_6": "value_6",
		},
	}

	dto := &dto.GeneratedData{
		Data: []map[string]any{obj1, obj2},
	}

	// Act
	resp := FromGeneratedDataDTOToGeneratedDataResponse(dto)

	// Assert
	assert.NotNil(t, resp)
	assert.Equal(t, len(dto.Data), len(resp.Data))
	assert.Equal(t, obj1["field_1"], resp.Data[0]["field_1"])
	assert.Equal(t, obj2["fields_5_6"], resp.Data[1]["fields_5_6"])
}
