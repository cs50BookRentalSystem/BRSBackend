package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/dto"
)

func (h *Handler) CreateRentTransaction(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.StudentID == uuid.Nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "student_id is required")
		return
	}

	if len(req.BookIDs) == 0 {
		h.writeErrorResponse(w, http.StatusBadRequest, "book_ids is required and cannot be empty")
		return
	}

	for _, bookID := range req.BookIDs {
		if bookID == uuid.Nil {
			h.writeErrorResponse(w, http.StatusBadRequest, "all book_ids must be valid UUIDs")
			return
		}
	}

	response, err := h.RentService.CreateRentTransaction(r.Context(), req)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to create rent transaction")
		return
	}

	h.writeResponse(w, http.StatusCreated, response)
}

func (h *Handler) ListRents(w http.ResponseWriter, r *http.Request, params api.ListRentsParams) {

	filter := dto.RentFilters{
		BookName:    params.BookName,
		StudentName: params.StudentName,
		Date:        &params.Date.Time,
		Offset:      int(*params.Offset),
		Limit:       int(*params.Limit),
	}

	rents, err := h.RentService.GetRents(r.Context(), filter)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to get rents")
		return
	}

	h.writeResponse(w, http.StatusOK, rents)
}

func (h *Handler) GetRentedBooksByStudent(w http.ResponseWriter, r *http.Request, params api.GetRentedBooksByStudentParams) {

	response, err := h.RentService.GetRentedBooksByStudent(r.Context(), params.StudentCardId)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to get rented books")
		return
	}

	h.writeResponse(w, http.StatusOK, response)
}

func (h *Handler) ReturnBooks(w http.ResponseWriter, r *http.Request) {

	var req struct {
		CartID uuid.UUID `json:"cart_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.CartID == uuid.Nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "cart_id is required")
		return
	}

	response, err := h.RentService.ReturnBooks(r.Context(), req.CartID)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to return books")
		return
	}

	h.writeResponse(w, http.StatusOK, response)
}
