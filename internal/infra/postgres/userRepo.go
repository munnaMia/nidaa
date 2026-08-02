package postgres

import (
	"database/sql"

	"github.com/munnaMia/nidaa/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

// create a user repository
func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}
