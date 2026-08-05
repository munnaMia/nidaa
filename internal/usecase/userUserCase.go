package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/munnaMia/nidaa/internal/domain"
)

type UserUseCase struct {
	repo         domain.UserRepository
	tokenService domain.TokenService
}

func NewUserUseCase(
	r domain.UserRepository,
	tk domain.TokenService,
) *UserUseCase {
	return &UserUseCase{
		repo:         r,
		tokenService: tk,
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

	// generate a jwt token
	jwt, err := uc.tokenService.GenerateToken(domain.Payload{
		Sub:   strconv.Itoa(user.ID),
		Name:  user.Name,
		Email: user.Email,
		IAT:   time.Now().Unix(),
		EXP:   time.Now().Add(time.Hour * 24).Unix(),
	})
	if err != nil {
		return "", nil, err
	}

	//return tokenjwt, user, err
	return jwt, user, err

}
