package postgres

import (
	"database/sql"

	"github.com/munnaMia/nidaa/internal/domain"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) domain.PostRepository {
	return &postRepository{
		db: db,
	}
}
