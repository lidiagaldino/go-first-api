package handler

import (
	"time"

	"github.com/lidiagaldino/go-first-api/schemas"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Login     string    `json:"login"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}
type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type UpdateUserResponse struct {
	Message string               `json:"message"`
	Data    schemas.UserResponse `json:"data"`
}

type CreateUserResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

type DeleteUserResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

type ShowUserResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

type ListUsersResponse struct {
	Message string         `json:"message"`
	Data    []UserResponse `json:"data"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
