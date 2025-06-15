package handlers

import (
	"encoding/json"
	"net/http"

	"BRSBackend/pkg/services"
)

type Handler struct {
	bookService    services.BookService
	authService    services.AuthService
	studentService services.StudentService
}

func NewHandler(svc *services.Service) *Handler {
	return &Handler{
		bookService:    svc.Book,
		authService:    svc.Auth,
		studentService: svc.Student,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (h *Handler) writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message, Code: statusCode})
}

func (h *Handler) writeResponse(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
