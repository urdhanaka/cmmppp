package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cmmppp/internal/api/v1/dto"
	"cmmppp/internal/repositories"
	"cmmppp/pkg/utils"
)

type UserHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	result, err := h.repo.GetUsersPaginated(r.Context(), utils.Pagination{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		http.Error(w, "failed to fetch users", http.StatusInternalServerError)
		return
	}

	resp := utils.NewPaginatedData(dto.NewUserResponses(result.Data), result.TotalCount, result.Page, result.PageSize)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
