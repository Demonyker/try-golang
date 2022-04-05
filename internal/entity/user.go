// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import (
	"database/sql"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		FirstName  string         `json:"firstName"`
		LastName   string         `json:"lastName"`
		MiddleName sql.NullString `json:"middleName"`
		Phone      string         `json:"phone"`
	}
)
