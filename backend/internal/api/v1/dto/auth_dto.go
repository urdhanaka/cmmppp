package dto

import "errors"

var (
	errEmailRequired    = errors.New("email is required")
	errPasswordRequired = errors.New("password is required")
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginRequest) Validate() error {
	if r.Email == "" {
		return errEmailRequired
	}
	if r.Password == "" {
		return errPasswordRequired
	}

	return nil
}
