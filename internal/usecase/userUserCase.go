package usecase

import "github.com/munnaMia/nidaa/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}
