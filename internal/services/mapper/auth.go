package mapper

import (
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

func FromUserDTOToUserEntity(dto *dto.UserDTO) *entities.User {
	return &entities.User{
		Id:        dto.Id,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

func FromUserEntityToUserDTO(entity *entities.User) *dto.UserDTO {
	return &dto.UserDTO{
		Id:        entity.Id,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
