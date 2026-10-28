package handlers

import (
	"cmmppp/internal/api/v1/dto"
	"cmmppp/internal/services"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

var errEmailExists = services.ErrEmailExists

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
	// page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	// pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	// result, err := h.service.GetUsersPaginated(r.Context(), utils.Pagination{
	// 	Page:     page,
	// 	PageSize: pageSize,
	// })
	// if err != nil {
	// 	slog.ErrorContext(r.Context(), "failed to fetch users", "err", err)
	// 	http.Error(w, "failed to fetch users", http.StatusInternalServerError)
	// 	return
	// }
	//
	// resp := utils.NewPaginatedData(dto.NewUserResponses(result.Data), result.TotalCount, result.Page, result.PageSize)
	//
	// w.Header().Set("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(resp); err != nil {
	// 	slog.ErrorContext(r.Context(), "failed to encode response", "err", err)
	// 	http.Error(w, "failed to encode response", http.StatusInternalServerError)
	// }
}

func (h *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.NewResponse(false, err.Error()))
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
