package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"cmmppp/internal/api/v1/dto"
	"cmmppp/internal/services"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{
		service,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to login", "err", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}

	response := dto.NewResponseWithData(true, "successfully logged in", dto.NewLoginResponse(token))

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "err", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
