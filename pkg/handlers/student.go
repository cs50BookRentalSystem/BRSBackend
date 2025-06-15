package handlers

import (
	"encoding/json"
	"net/http"

	oapiTypes "github.com/oapi-codegen/runtime/types"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
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

	if err := h.studentService.CreateStudent(r.Context(), &student); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	h.writeResponse(w, http.StatusCreated, map[string]string{"message": "Student added successfully"})
}

func (h *Handler) GetStudentById(w http.ResponseWriter, r *http.Request, id oapiTypes.UUID) {

	student, err := h.studentService.GetStudentByID(r.Context(), id.String())
	if err != nil {
		return
	}

	h.writeResponse(w, http.StatusOK, student)
}
