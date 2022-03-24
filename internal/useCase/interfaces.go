package useCase

import (
	"context"

	"fairseller-backend/internal/entity"
)

type (
	// Auth - auth useCase.
	Auth interface {
		SignUpRequest(context.Context, entity.SignUpRequestDto) error
	}

	// UserRepository - user repository.
	UserRepository interface {
		Store(context.Context, entity.User) error
	}
)
