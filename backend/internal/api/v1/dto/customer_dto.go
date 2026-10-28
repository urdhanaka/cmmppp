package dto

import (
	"cmmppp/internal/entities"
	"errors"
	"time"

	"github.com/google/uuid"
)

var errCustomerNameRequired = errors.New("customer_name is required")

type (
	CreateCustomerRequest struct {
		CustomerName string `json:"customer_name"`
	}

	CustomerResponse struct {
		ID           uuid.UUID `json:"id"`
		CustomerName string    `json:"customer_name"`
		CreatedAt    time.Time `json:"created_at"`
	}
)

func (r CreateCustomerRequest) Validate() error {
	if r.CustomerName == "" {
		return errCustomerNameRequired
	}

	return nil
}

func (r CreateCustomerRequest) ToEntity() entities.Customer {
	return entities.Customer{
		CustomerName: r.CustomerName,
	}
}

func NewCustomerResponse(c entities.Customer) CustomerResponse {
	return CustomerResponse{
		ID:           c.ID,
		CustomerName: c.CustomerName,
		CreatedAt:    c.CreatedAt,
	}
}

func NewCustomerResponses(customers []entities.Customer) []CustomerResponse {
	resp := make([]CustomerResponse, len(customers))

	for i, c := range customers {
		resp[i] = NewCustomerResponse(c)
	}

	return resp
}
