package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
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

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	query := `
		INSERT INTO users(name, username, email, password_hash, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query, u.Name, u.UserName, u.Email, u.Password, u.CreatedAt).Scan(&u.ID)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			details := strings.ToLower(pgErr.Detail)
			constraint := strings.ToLower(pgErr.ConstraintName)

			switch {
			case strings.Contains(constraint, "username"), strings.Contains(details, "username"):
				return domain.ErrUsernameAlreadyExist
			case strings.Contains(constraint, "email"), strings.Contains(details, "email"):
				return domain.ErrEmailAlreadyExist
			default:
				return domain.ErrEmailAlreadyExist
			}
		}
		return fmt.Errorf("userRepository err: %w", err)
	}

	return nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT  id, name, username, email, password_hash, created_at, updated_at 
		FROM users
		WHERE email = $1
	`

	var user domain.User

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("failed to fetch user by email: %w", err)
	}

	return &user, nil
}
