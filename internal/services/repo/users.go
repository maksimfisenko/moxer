package repo

import (
	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
)

type UsersRepo interface {
	Create(user *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindById(id uuid.UUID) (*entities.User, error)
}
