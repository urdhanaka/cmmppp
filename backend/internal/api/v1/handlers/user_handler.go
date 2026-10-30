package handlers

import (
	"cmmppp/internal/api/v1/dto"
	"cmmppp/internal/services"
	"cmmppp/pkg/utils"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	errEmailExists  = services.ErrEmailExists
	errInvalidID    = services.ErrInvalidID
	errUserNotFound = services.ErrUserNotFound
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary		Get all users without pagination
// @Description	Get all users without pagination
// @Tags		users
// @Produce		json
// @Success		200	{array}	dto.UserResponse
// @Router		/users [get]
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	w.Header().Set("Content-Type", "application/json")

	result, err := h.service.GetUsersPaginated(r.Context(), &utils.PaginationMeta{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to fetch users", "err", err)
		http.Error(w, "failed to fetch users", http.StatusInternalServerError)
		return
	}

	// map the result to the dto
	data := utils.PaginatedData[dto.UserResponse]{
		Items: dto.NewUserResponses(&result.Items),
		Meta:  result.Meta,
	}

	response := dto.NewResponseWithData(true, "successfully retrieved all users data", data)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "err", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.NewResponse(false, "invalid request body"))
		return
	}

	createdUser, err := h.service.CreateUser(r.Context(), &services.CreateUserInput{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		if errors.Is(err, errEmailExists) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(dto.NewResponse(false, errEmailExists.Error()))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.NewResponse(false, "failed to create user"))
		return
	}

	response := dto.NewResponseWithData(true, "User succesfully created", dto.NewUserResponse(createdUser))

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "err", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	w.Header().Set("Content-Type", "application/json")

	err := h.service.DeleteUser(r.Context(), userID)
	if err != nil {
		if errors.Is(err, errInvalidID) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(dto.NewResponse(false, errInvalidID.Error()))
			return
		}

		if errors.Is(err, errUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(dto.NewResponse(false, errUserNotFound.Error()))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.NewResponse(false, "failed to delete user"))
		return
	}

	response := dto.NewResponse(true, "successfully deleted the user")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "err", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
