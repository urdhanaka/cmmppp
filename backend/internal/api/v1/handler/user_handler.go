package handler

import (
	"cmmppp/internal/api/v1/dto"
	"cmmppp/internal/repositories"
	"cmmppp/pkg/utils"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type UserHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
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

	result, err := h.repo.GetUsersPaginated(r.Context(), utils.Pagination{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to fetch users", "err", err)
		http.Error(w, "failed to fetch users", http.StatusInternalServerError)
		return
	}

	resp := utils.NewPaginatedData(dto.NewUserResponses(result.Data), result.TotalCount, result.Page, result.PageSize)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.ErrorContext(r.Context(), "failed to encode response", "err", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
