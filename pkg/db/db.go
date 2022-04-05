package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
