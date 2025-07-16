package db

import (
	"github.com/maksimfisenko/moxer/internal/config"
	"github.com/maksimfisenko/moxer/internal/repo/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := config.Cfg.DB.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.User{},
		&entities.Template{},
	)
}
