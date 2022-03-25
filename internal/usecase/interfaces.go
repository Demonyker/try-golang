package usecase

import (
	"context"

	"fairseller-backend/internal/entity"
)

type (
	// Auth - auth usecase.
	Auth interface {
		SignUpRequest(context.Context, entity.SignUpRequest) error
	}

	// UserRepository - user repository.
	UserRepository interface {
		Store(context.Context, entity.User) error
		GetOneByPhone(context.Context, string) (entity.User, error)
	}
)
