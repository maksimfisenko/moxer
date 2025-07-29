package mapper

import (
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

func FromUserDTOToUserEntity(dto *dto.UserDTO, hash string) *entities.User {
	return &entities.User{
		Id:           dto.Id,
		Email:        dto.Email,
		PasswordHash: hash,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
	}
}

func FromUserEntityToUserDTO(entity *entities.User) *dto.UserDTO {
	return &dto.UserDTO{
		Id:        entity.Id,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
