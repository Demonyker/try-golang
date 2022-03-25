package repository

import (
	"context"
	"fairseller-backend/internal/entity"
	"fairseller-backend/pkg/postgres"
	"fmt"
)

// UserRepository -.
type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) *UserRepository {
	return &UserRepository{pg}
}

// Store -.
func (r *UserRepository) Store(ctx context.Context, userData entity.User) error {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("first_name, last_name, middle_name, phone").
		Values(userData.FirstName, userData.LastName, userData.MiddleName, userData.Phone).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepository - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepository - Store - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *UserRepository) GetOneByPhone(ctx context.Context, phone string) (entity.User, error) {
	user := entity.User{}
	sql, args, err := r.Builder.
		Select("id, first_name, last_name, middle_name, phone").
		From("users").
		Where("users.phone = ?", phone).
		ToSql()
	if err != nil {
		return user, fmt.Errorf("UserRepository - GetOneByPhone - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return user, fmt.Errorf("UserRepository - GetOneByPhone - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.MiddleName, &user.Phone)
		if err != nil {
			return user, fmt.Errorf("UserRepository - GetOneByPhone - rows.Scan: %w", err)
		}
	}

	return user, nil
}
