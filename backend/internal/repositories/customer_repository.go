package repositories

import (
	"cmmppp/internal/entities"
	"cmmppp/pkg/utils"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrCustomerNotFound = errors.New("customer not found")

type (
	CustomerRepository interface {
		PostCustomer(ctx context.Context, customer *entities.Customer) error
		GetAllCustomers(ctx context.Context) ([]entities.Customer, error)
		GetCustomersPaginated(ctx context.Context, p utils.Pagination) (utils.PaginatedData[entities.Customer], error)
		DeleteCustomer(ctx context.Context, customer *entities.Customer) error
	}

	customerRepository struct {
		db *gorm.DB
	}
)

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db,
	}
}

func (r *customerRepository) PostCustomer(ctx context.Context, customer *entities.Customer) error {
	return gorm.G[entities.Customer](r.db).Create(ctx, customer)
}

func (r *customerRepository) GetCustomersPaginated(
	ctx context.Context,
	p utils.Pagination,
) (utils.PaginatedData[entities.Customer], error) {
	p.Normalize()

	total, err := gorm.G[entities.Customer](r.db).Count(ctx, "*")
	if err != nil {
		return utils.PaginatedData[entities.Customer]{}, err
	}

	customers, err := gorm.G[entities.Customer](r.db).
		Order("created_at").
		Limit(p.PageSize).
		Offset((p.Page - 1) * p.PageSize).
		Find(ctx)
	if err != nil {
		return utils.PaginatedData[entities.Customer]{}, err
	}

	return utils.NewPaginatedData(customers, total, p.Page, p.PageSize), nil
}

func (r *customerRepository) GetAllCustomers(ctx context.Context) ([]entities.Customer, error) {
	return gorm.G[entities.Customer](r.db).Find(ctx)
}

func (r *customerRepository) DeleteCustomer(ctx context.Context, customer *entities.Customer) error {
	rows, err := gorm.G[entities.Customer](r.db).Where(customer).Delete(ctx)
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrCustomerNotFound
	}

	return nil
}
