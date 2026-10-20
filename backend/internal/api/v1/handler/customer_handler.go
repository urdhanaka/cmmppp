package handler

import (
	"cmmppp/internal/api/v1/dto"
	"cmmppp/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.service.GetAllCustomers(r.Context())
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to fetch customers", "err", err)
		http.Error(w, "failed to fetch customers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.NewCustomerResponses(customers)); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "err", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := h.service.CreateCustomer(r.Context(), req.ToEntity())
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to create customer", "err", err)
		http.Error(w, "failed to create customer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(dto.NewCustomerResponse(customer)); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "err", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
