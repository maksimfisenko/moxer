package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/handlers/requests"
	"github.com/maksimfisenko/moxer/internal/handlers/responses"
	"github.com/maksimfisenko/moxer/internal/services/dto"
)

func FromCredentialsRequestToUserDTO(cred *requests.CredentialsRequest) *dto.User {
	return &dto.User{
		Id:        uuid.New(),
		Email:     cred.Email,
		Password:  cred.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func FromUserDTOToUserResponse(dto *dto.User) *responses.UserResponse {
	return &responses.UserResponse{
		Id:        dto.Id.String(),
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

func FromLoginRequestToUserCredentialsDTO(req *requests.CredentialsRequest) *dto.Credentials {
	return &dto.Credentials{
		Email:    req.Email,
		Password: req.Password,
	}
}

func FromTokenDTOToTokenResponse(dto *dto.Token) *responses.Token {
	return &responses.Token{
		Token: dto.Token,
	}
}
