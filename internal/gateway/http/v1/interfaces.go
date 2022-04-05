package v1

import (
	"context"
)

type (
	AuthUseCase interface {
		SignUpRequest(ctx context.Context, phone string) error
	}
)
