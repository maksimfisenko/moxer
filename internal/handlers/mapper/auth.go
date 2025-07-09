package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/handlers/requests"
	"github.com/maksimfisenko/moxer/internal/handlers/responses"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

func FromRegisterRequestToUserDTO(req *requests.RegisterRequest) *dto.UserDTO {
	return &dto.UserDTO{
		Id:        uuid.New(),
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func FromUserDTOToUserResponse(dto *dto.UserDTO) *responses.UserResponse {
	return &responses.UserResponse{
		Id:        dto.Id.String(),
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
