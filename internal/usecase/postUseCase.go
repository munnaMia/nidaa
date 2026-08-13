package usecase

import "github.com/munnaMia/nidaa/internal/domain"

type PostUseCase struct {
	repo domain.PostRepository
}

func NewPostUseCase(r domain.PostRepository) *PostUseCase {
	return &PostUseCase{
		repo: r,
	}
}
