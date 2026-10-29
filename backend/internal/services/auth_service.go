package services

import (
	"context"
	"errors"

	"cmmppp/internal/config"
	"cmmppp/internal/repositories"
	"cmmppp/pkg/utils"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type (
	AuthService interface {
		Login(ctx context.Context, email, password string) (string, error)
	}

	authService struct {
		userRepo     repositories.UserRepository
		tokenManager *config.TokenManager
	}
)

func NewAuthService(userRepo repositories.UserRepository, tokenManager *config.TokenManager) AuthService {
	return &authService{
		userRepo:     userRepo,
		tokenManager: tokenManager,
	}
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}

		return "", err
	}

	if !utils.ComparePassword(user.Password, password) {
		return "", ErrInvalidCredentials
	}

	role := "user"
	if user.IsPlatformAdmin {
		role = "platform_admin"
	}

	return s.tokenManager.GenerateToken(user.ID, role)
}
