package repository

import (
	"context"
	"fmt"

	"fairseller-backend/internal/entity"
	"fairseller-backend/pkg/postgres"
)

// userRepository -.
type userRepository struct {
	*postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *userRepository {
	return &userRepository{pg}
}

// Store -.
func (r *userRepository) Store(ctx context.Context, userData entity.User) error {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("first_name, last_name, middle_name, phone").
		Values(userData.FirstName, userData.LastName, userData.MiddleName, userData.Phone).
		ToSql()
	if err != nil {
		return fmt.Errorf("userRepository - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("userRepository - Store - r.Pool.Exec: %w", err)
	}

	return nil
}
