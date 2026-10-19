package service

import (
	"cmmppp/internal/entities"
	"cmmppp/internal/repositories"
	"context"
)

type (
	CustomerService interface {
		GetAllCustomers(ctx context.Context) ([]entities.Customer, error)
		CreateCustomer(ctx context.Context, customer entities.Customer) (entities.Customer, error)
	}

	customerService struct {
		repo repositories.CustomerRepository
	}
)

func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerService{
		repo,
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
