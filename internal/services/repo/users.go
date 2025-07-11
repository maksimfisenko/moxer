package repo

import "github.com/maksimfisenko/moxer/internal/repo/entities"

type UsersRepo interface {
	Create(user *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}
