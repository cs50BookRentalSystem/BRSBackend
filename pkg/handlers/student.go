package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	oapiTypes "github.com/oapi-codegen/runtime/types"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	"BRSBackend/pkg/validation"
)

func (h *Handler) ListAllStudents(w http.ResponseWriter, r *http.Request, params api.ListAllStudentsParams) {

	paginationParams := dto.PaginationParams{
		Limit:  10,
		Offset: 0,
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if params.Limit != nil && int(*params.Limit) > 0 {
			paginationParams.Limit = int(*params.Limit)
		} else {
			h.writeErrorResponse(w, http.StatusBadRequest, "Invalid limit parameter")
			return
		}
	}

	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if params.Offset != nil && int(*params.Limit) > 0 {
			paginationParams.Offset = int(*params.Offset)
		} else {
			h.writeErrorResponse(w, http.StatusBadRequest, "Invalid offset parameter")
		}
	}

	if cardId := r.URL.Query().Get("card_id"); cardId != "" {
		if params.CardId != nil && strings.TrimSpace(*params.CardId) != "" {
			response, err := h.studentService.GetStudentByCardNumber(r.Context(), cardId)
			if err != nil {
				h.writeErrorResponse(w, http.StatusNotFound, "Failed to get student")
				return
			}
			h.writeResponse(w, http.StatusOK, response)
			return
		} else {
			h.writeErrorResponse(w, http.StatusBadRequest, "Invalid offset parameter")
		}
	}

	allStudents, err := h.studentService.GetAllStudents(r.Context(), paginationParams)
	if err != nil {
		return
	}

	students := make([]api.Students, len(allStudents.Results))
	for i, student := range allStudents.Results {
		students[i] = *student
	}

	apiPagination := &api.PaginationInfo{
		Offset:      &allStudents.Pagination.Offset,
		Limit:       &allStudents.Pagination.Limit,
		Total:       &allStudents.Pagination.Total,
		HasNext:     &allStudents.Pagination.HasNext,
		HasPrevious: &allStudents.Pagination.HasPrevious,
	}

	h.writeResponse(w, http.StatusOK, api.ListAllStudents200JSONResponse{Results: &students, Pagination: apiPagination})
}

func (h *Handler) AddStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if validationErrors := validation.ValidateStruct(student); validationErrors != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, validation.FormatErrors(validationErrors))
		return
	}

	if err := h.studentService.CreateStudent(r.Context(), &student); err != nil {
		h.writeErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	h.writeResponse(w, http.StatusCreated, map[string]string{"message": "Student added successfully"})
}

func (h *Handler) GetStudentById(w http.ResponseWriter, r *http.Request, id oapiTypes.UUID) {

	if id == uuid.Nil || id.String() == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "Student ID is required")
		return
	}
	student, err := h.studentService.GetStudentByID(r.Context(), id.String())
	if err != nil {
		h.writeErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	h.writeResponse(w, http.StatusOK, student)
}

func (h *Handler) DeleteStudentById(w http.ResponseWriter, r *http.Request, id oapiTypes.UUID) {
	if id == uuid.Nil || id.String() == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "Student ID is required")
		return
	}
	if err := h.studentService.DeleteStudent(r.Context(), id.String()); err != nil {
		h.writeErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	h.writeResponse(w, http.StatusCreated, "ok")
}
