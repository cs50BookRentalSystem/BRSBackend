package dto

import "github.com/google/uuid"

type LoginRequest struct {
	User string `json:"user" validate:"required"`
	Pass string `json:"pass" validate:"required"`
}

type LoginResponse struct {
	Message     string    `json:"message"`
	LibrarianId uuid.UUID `json:"librarian_id"`
}
