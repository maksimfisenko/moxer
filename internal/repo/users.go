package repo

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/errorsx"
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
	err := ur.db.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, errorsx.ErrEmailAlreadyExists
		}
		return nil, err
	}

	return ur.FindById(user.Id)
}

func (ur *usersRepo) FindById(id uuid.UUID) (*entities.User, error) {
	var user entities.User

	err := ur.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *usersRepo) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
