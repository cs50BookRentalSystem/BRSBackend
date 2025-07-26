package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/validation"
)

func (h *Handler) CreateRentTransaction(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if validationErrors := validation.ValidateStruct(req); validationErrors != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, validation.FormatErrors(validationErrors))
		return
	}

	response, err := h.rentService.CreateRentTransaction(r.Context(), req)
	if err != nil {
		h.writeErrorResponse(w, http.StatusUnprocessableEntity, "Failed to create rent transaction")
		return
	}

	h.writeResponse(w, http.StatusCreated, response)
}

func (h *Handler) ListRents(w http.ResponseWriter, r *http.Request, params api.ListRentsParams) {
	defaultBook := ""
	defaultStudent := ""
	defaultDate := time.Now()
	filter := dto.RentFilters{
		BookName:    &defaultBook,
		StudentName: &defaultStudent,
		Date:        &defaultDate,
		Limit:       10,
		Offset:      0,
	}

	if params.BookName != nil && strings.TrimSpace(*params.BookName) != "" {
		filter.BookName = params.BookName
	}

	if params.StudentName != nil && strings.TrimSpace(*params.StudentName) != "" {
		filter.StudentName = params.StudentName
	}
	if params.Date != nil {
		filter.Date = &params.Date.Time
	}

	if params.Limit != nil && int(*params.Limit) > 0 {
		filter.Limit = int(*params.Limit)
	}

	if params.Offset != nil && int(*params.Offset) > 0 {
		filter.Offset = int(*params.Offset)
	}

	rents, err := h.rentService.GetRents(r.Context(), filter)
	if err != nil {
		h.writeErrorResponse(w, http.StatusUnprocessableEntity, "Failed to get rents")
		return
	}

	h.writeResponse(w, http.StatusOK, rents)
}

func (h *Handler) GetRentedBooksByStudent(w http.ResponseWriter, r *http.Request, params api.GetRentedBooksByStudentParams) {

	if params.StudentCardId == nil || strings.TrimSpace(*params.StudentCardId) == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "Student card ID is required")
		return
	}

	response, err := h.rentService.GetRentedBooksByStudent(r.Context(), params.StudentCardId)
	if err != nil {
		h.writeErrorResponse(w, http.StatusUnprocessableEntity, "Failed to get rented books")
		return
	}

	h.writeResponse(w, http.StatusOK, response)
}

func (h *Handler) ReturnBooks(w http.ResponseWriter, r *http.Request) {

	var req struct {
		CartID uuid.UUID `json:"cart_id" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if validationErrors := validation.ValidateStruct(req); validationErrors != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, validation.FormatErrors(validationErrors))
		return
	}

	response, err := h.rentService.ReturnBooks(r.Context(), req.CartID)
	if err != nil {
		h.writeErrorResponse(w, http.StatusUnprocessableEntity, "Failed to return books")
		return
	}

	h.writeResponse(w, http.StatusOK, response)
}
