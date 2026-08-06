package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrEmailAlreadyExist    = errors.New("email already register")
	ErrUsernameAlreadyExist = errors.New("username already taken")
	ErrInvalidCredentials   = errors.New("Invalid email or password")
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	UserName  string    `json:"userName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}
