package dto

import (
	"cmmppp/internal/entities"
	"time"

	"github.com/google/uuid"
)

// UserResponse is the HTTP-facing shape of a user — deliberately excludes
// Password and other internal-only fields from entities.User.
type (
	UserResponse struct {
		ID              uuid.UUID `json:"id"`
		Email           string    `json:"email"`
		IsPlatformAdmin bool      `json:"is_platform_admin"`
		CreatedAt       time.Time `json:"created_at"`
	}

	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
)

func NewUserResponse(u *entities.User) UserResponse {
	return UserResponse{
		ID:              u.ID,
		Email:           u.Email,
		IsPlatformAdmin: u.IsPlatformAdmin,
		CreatedAt:       u.CreatedAt,
	}
}

func NewUserResponses(users []entities.User) []UserResponse {
	resp := make([]UserResponse, len(users))

	for i, u := range users {
		resp[i] = NewUserResponse(&u)
	}

	return resp
}
