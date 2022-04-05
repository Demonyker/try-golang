package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository *UserRepository
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
	}
}
