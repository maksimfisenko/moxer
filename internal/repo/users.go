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

func (ur *usersRepo) Create(user *entities.User) (*entities.User, error) {
	if err := ur.db.Create(user).Error; err != nil {
		return nil, err
	}

	return ur.FindById(user.Id)
}

func (ur *usersRepo) FindById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *usersRepo) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
