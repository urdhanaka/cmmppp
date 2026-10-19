package repositories

import (
	"cmmppp/internal/entities"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrCustomerNotFound = errors.New("customer not found")

type (
	CustomerRepository interface {
		PostCustomer(ctx context.Context, customer *entities.Customer) error
		GetAllCustomers(ctx context.Context) ([]entities.Customer, error)
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
