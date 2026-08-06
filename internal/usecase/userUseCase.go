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

func (uc *UserUseCase) LoginUser(ctx context.Context, email, password string) (string, *domain.User, error) {
	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, err
	}

	// compare user give password with my database hash
	password_hash := password // create a password hash to compare
	if password_hash != user.Password {
		return "", nil, domain.ErrInvalidCredentials
	}

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

	return jwt, user, nil
}
