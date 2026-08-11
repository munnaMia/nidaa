package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/munnaMia/nidaa/internal/domain"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassBytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", nil, err
	}

	user := &domain.User{
		UserName:  username,
		Name:      name,
		Email:     email,
		Password:  string(hashedPassBytes),
		CreatedAt: time.Now().UTC(),
	}

	err = uc.repo.Create(ctx, user)
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

	//return tokenjwt, user
	return jwt, user, nil
}

func (uc *UserUseCase) LoginUser(ctx context.Context, email, password string) (string, *domain.User, error) {
	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", nil, err
	}

	// Compare incoming plain password against stored bcrypt hash
	// bcrypt handles constant-time comparison internally!
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
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

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (*domain.User, error) {
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalide user id: %w", err)
	}

	user, err := uc.repo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
