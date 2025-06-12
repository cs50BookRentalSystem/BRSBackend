package handlers

import (
	_ "context"
	"encoding/json"
	"net/http"

	oapiTypes "github.com/oapi-codegen/runtime/types"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	_ "BRSBackend/pkg/repository"
	"BRSBackend/pkg/services"
)

type BookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

//type APIResponse struct {
//	Success bool        `json:"success"`
//	Data    interface{} `json:"data,omitempty"`
//	Error   string      `json:"error,omitempty"`
//	Message string      `json:"message,omitempty"`
//}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (h *BookHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message, Code: statusCode})
}

func (h *BookHandler) writeResponse(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

//func (h *BookHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
//	h.writeResponse(w, statusCode, APIResponse{
//		Success: false,
//		Error:   message,
//	})
//}
//
//func (h *BookHandler) writeSuccessResponse(w http.ResponseWriter, data interface{}, message string) {
//	h.writeResponse(w, http.StatusOK, APIResponse{
//		Success: true,
//		Data:    data,
//		Message: message,
//	})
//}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.bookService.CreateBook(r.Context(), &book); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	h.writeResponse(w, http.StatusCreated, map[string]string{"message": "Book added successfully"})

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) ListOrSearchBooks(w http.ResponseWriter, r *http.Request, params api.ListOrSearchBooksParams) {
	paginationParams := dto.PaginationParams{
		Query:  "",
		Limit:  10,
		Offset: 0,
	}
	if params.Query != nil {
		paginationParams.Query = *params.Query
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

	response, err := h.bookService.GetAllBooks(r.Context(), paginationParams)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	apiBooks := make([]api.Books, len(response.Results))
	for i, book := range response.Results {
		apiBooks[i] = *book
	}

	apiPagination := &api.PaginationInfo{
		Offset:      &response.Pagination.Offset,
		Limit:       &response.Pagination.Limit,
		Total:       &response.Pagination.Total,
		HasNext:     &response.Pagination.HasNext,
		HasPrevious: &response.Pagination.HasPrevious,
	}

	h.writeResponse(w, http.StatusOK, api.ListOrSearchBooks200JSONResponse{Results: &apiBooks, Pagination: apiPagination})

}

func (h *BookHandler) Login(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *BookHandler) Logout(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *BookHandler) ListOverdueRentals(w http.ResponseWriter, r *http.Request, params api.ListOverdueRentalsParams) {
	panic("implement me")
}

func (h *BookHandler) ListRents(w http.ResponseWriter, r *http.Request, params api.ListRentsParams) {
	panic("implement me")
}

func (h *BookHandler) CreateRentTransaction(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *BookHandler) GetRentalReports(w http.ResponseWriter, r *http.Request, params api.GetRentalReportsParams) {
	panic("implement me")
}

func (h *BookHandler) GetRentedBooksByStudent(w http.ResponseWriter, r *http.Request, params api.GetRentedBooksByStudentParams) {
	panic("implement me")
}

func (h *BookHandler) ReturnBooks(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *BookHandler) ListAllStudents(w http.ResponseWriter, r *http.Request, params api.ListAllStudentsParams) {
	panic("implement me")
}

func (h *BookHandler) AddStudent(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *BookHandler) GetStudentById(w http.ResponseWriter, r *http.Request, id oapiTypes.UUID) {
	panic("implement me")
}
