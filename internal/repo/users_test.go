package repo

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDB() (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("falied to connect to sqlite db")
	}

	err = db.AutoMigrate(
		&entities.User{},
		&entities.Template{},
	)
	if err != nil {
		panic("failed to migrate db")
	}

	cleanup := func() {
		db.Exec("DELETE FROM users")
	}

	return db, cleanup
}

func TestUsersRepo_Create(t *testing.T) {
	// Arange
	db, cleanup := setupDB()
	defer cleanup()

	usersRepo := NewUsersRepo(db)

	user := &entities.User{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "11111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Act
	createdUser, err := usersRepo.Create(user)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, user.Id, createdUser.Id)
	assert.Equal(t, user.Email, createdUser.Email)
}

func TestUsersRepo_FindById(t *testing.T) {
	// Arange
	db, cleanup := setupDB()
	defer cleanup()

	usersRepo := NewUsersRepo(db)

	user := &entities.User{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "11111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = usersRepo.Create(user)

	// Act
	fetchedUser, err := usersRepo.FindById(user.Id)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, user.Email, fetchedUser.Email)
}

func TestUsersRepo_FindByEmail(t *testing.T) {
	// Arange
	db, cleanup := setupDB()
	defer cleanup()

	usersRepo := NewUsersRepo(db)

	user := &entities.User{
		Id:        uuid.New(),
		Email:     "test@example.com",
		Password:  "11111111",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = usersRepo.Create(user)

	// Act
	fetchedUser, err := usersRepo.FindByEmail(user.Email)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, user.Id, fetchedUser.Id)
}
