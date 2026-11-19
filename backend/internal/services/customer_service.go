package services

import (
	"cmmppp/internal/entities"
	"cmmppp/internal/repositories"
	"cmmppp/internal/utils"
	"context"
)

type (
	CustomerService interface {
		GetAllCustomers(ctx context.Context) ([]entities.Customer, error)
		CreateCustomer(ctx context.Context, customer entities.Customer) (entities.Customer, error)
		GetCustomersPaginated(ctx context.Context, p *utils.PaginationMeta) (utils.PaginatedData[entities.Customer], error)
	}

	customerService struct {
		repo repositories.CustomerRepository
	}
)

func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerService{
		repo: repo,
	}
}

func (s *customerService) GetAllCustomers(ctx context.Context) ([]entities.Customer, error) {
	return s.repo.GetAllCustomers(ctx)
}

func (s *customerService) CreateCustomer(ctx context.Context, customer entities.Customer) (entities.Customer, error) {
	if err := s.repo.PostCustomer(ctx, &customer); err != nil {
		return entities.Customer{}, err
	}

	return customer, nil
}

func (s *customerService) GetCustomersPaginated(
	ctx context.Context,
	p *utils.PaginationMeta,
) (utils.PaginatedData[entities.Customer], error) {
	p.Normalize()

	return s.repo.GetCustomersPaginated(ctx, p)
}
