package services

import (
	"cmmppp/internal/entities"
	"cmmppp/internal/repositories"
	"cmmppp/pkg/utils"
	"context"
	"errors"
)

var ErrEmailExists = repositories.ErrEmailExists

type (
	UserService interface {
		CreateUser(ctx context.Context, userInput *CreateUserInput) (*entities.User, error)
	}

	userService struct {
		repo repositories.UserRepository
	}
)

type (
	CreateUserInput struct {
		Email    string
		Name     string
		Password string
	}
)

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, userInput *CreateUserInput) (*entities.User, error) {
	hashedPassword, err := utils.HashPassword(userInput.Password)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.User{
		Email:    userInput.Email,
		Name:     userInput.Name,
		Password: hashedPassword,
	}

	if err := s.repo.PostUser(ctx, userEntity); err != nil {
		if errors.Is(err, ErrEmailExists) {
			return nil, ErrEmailExists
		}

		return nil, err
	}

	return userEntity, nil
}
