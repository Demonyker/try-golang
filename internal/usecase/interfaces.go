package usecase

import (
	"fairseller-backend/internal/entity"
)

type (
	userRepositoryInterface interface {
		GetOneByPhone(phone string) (entity.User, error)
	}
)
