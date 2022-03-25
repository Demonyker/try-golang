package repository

import "fairseller-backend/pkg/postgres"

// const _defaultEntityCap = 64

type Repository struct {
	UserRepository *UserRepository
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(pg),
	}
}
