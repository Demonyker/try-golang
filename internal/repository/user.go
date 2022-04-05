package repository

import (
	"fairseller-backend/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetOneByPhone(phone string) (entity.User, error) {
	user := entity.User{}

	result := r.DB.Where("phone = ?", phone).Find(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
