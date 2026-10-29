package repositories

import (
	"cmmppp/internal/entities"
	"cmmppp/pkg/utils"
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailExists  = errors.New("email already exists")
)

type (
	UserRepository interface {
		PostUser(ctx context.Context, user *entities.User) error
		GetAllUsers(ctx context.Context) ([]entities.User, error)
		GetUsersPaginated(ctx context.Context, p *utils.PaginationMeta) (utils.PaginatedData[entities.User], error)
		GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
		DeleteUser(ctx context.Context, user *entities.User) error
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	return gorm.G[entities.User](r.db).Find(ctx)
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := gorm.G[entities.User](r.db).Where("email = ?", email).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) PostUser(ctx context.Context, user *entities.User) error {
	if err := gorm.G[entities.User](r.db).Create(ctx, user); err != nil {
		return errorMapper(err)
	}

	return nil
}

func (r *userRepository) GetUsersPaginated(
	ctx context.Context,
	p *utils.PaginationMeta,
) (utils.PaginatedData[entities.User], error) {
	total, err := gorm.G[entities.User](r.db).Count(ctx, "*")
	if err != nil {
		return utils.PaginatedData[entities.User]{}, err
	}

	users, err := gorm.G[entities.User](r.db).
		Order("created_at").
		Limit(p.PageSize).
		Offset((p.Page - 1) * p.PageSize).
		Find(ctx)
	if err != nil {
		return utils.PaginatedData[entities.User]{}, err
	}

	return utils.NewPaginatedData(users, total, p.Page, p.PageSize), nil
}

func (r *userRepository) DeleteUser(ctx context.Context, user *entities.User) error {
	rowsAffected, err := gorm.G[entities.User](r.db).Where(user).Delete(ctx)
	if rowsAffected == 0 {
		return ErrUserNotFound
	}
	if err != nil {
		return err
	}

	return nil
}
