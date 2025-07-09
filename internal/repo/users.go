package repo

import (
	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"gorm.io/gorm"
)

type usersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *usersRepo {
	return &usersRepo{db: db}
}

func (ur *usersRepo) findById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *usersRepo) Create(user *entities.User) (*entities.User, error) {
	if err := ur.db.Create(user).Error; err != nil {
		return nil, err
	}

	return ur.findById(user.Id)
}
