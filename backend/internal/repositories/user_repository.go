package repositories

import (
	"context"

	"cmmppp/internal/entities"
	"cmmppp/pkg/utils"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetAllUsers(ctx context.Context) ([]entities.User, error)
		GetUsersPaginated(ctx context.Context, p utils.Pagination) (utils.PaginatedData[entities.User], error)
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
