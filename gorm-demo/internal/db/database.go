package db

import (
	"gorm-demo/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *config.Config) (*gorm.DB, error) {

	dsn := config.Database.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
