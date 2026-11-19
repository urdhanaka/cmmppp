package services

import (
	"cmmppp/internal/repositories"
	"cmmppp/internal/utils"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
)

type (
	AuthService interface {
		Login(ctx context.Context, email, password string) (string, error)
	}

	TokenGenerator interface {
		GenerateToken(userID uuid.UUID, role string) (string, error)
	}

	authService struct {
		userRepo     repositories.UserRepository
		tokenManager TokenGenerator
	}
)

func NewAuthService(userRepo repositories.UserRepository, tokenManager TokenGenerator) AuthService {
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

	if !utils.ComparePassword(user.PasswordHash, password) {
		return "", ErrInvalidCredentials
	}

	role := "user"
	if user.IsPlatformAdmin {
		role = "platform_admin"
	}

	return s.tokenManager.GenerateToken(user.ID, role)
}
