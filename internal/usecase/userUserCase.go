package usecase

import (
	"context"
	"time"

	"github.com/munnaMia/nidaa/internal/domain"
)

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) RegisterUser(ctx context.Context, username, name, email, password string) (string, *domain.User, error) {
	user := &domain.User{
		UserName:  username,
		Name:      name,
		Email:     email,
		Password:  password, // hash the pssword latter.
		CreatedAt: time.Now().UTC(),
	}

	err := uc.repo.Create(ctx, user)
	if err != nil {
		return "", nil, err
	}

	//return tokenjwt, user, err
	return "", user, err

}
