package repositories

import (
	"cmmppp/internal/entities"
	"cmmppp/pkg/utils"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type (
	UserRepository interface {
		GetAllUsers(ctx context.Context) ([]entities.User, error)
		GetUsersPaginated(ctx context.Context, p utils.Pagination) (utils.PaginatedData[entities.User], error)
		GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
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

func (r *userRepository) GetUsersPaginated(
	ctx context.Context,
	p utils.Pagination,
) (utils.PaginatedData[entities.User], error) {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = utils.DEFAULT_PAGE_SIZE
	}

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

	return utils.PaginatedData[entities.User]{
		Data:       users,
		TotalCount: total,
		Page:       p.Page,
		PageSize:   p.PageSize,
	}, nil
}
