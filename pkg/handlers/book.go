package handlers

import (
	_ "context"
	"encoding/json"
	"net/http"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
)

//type BookHandler struct {
//	bookService services.BookService
//}
//
//func NewBookHandler(bookService services.BookService) *BookHandler {
//	return &BookHandler{
//		bookService: bookService,
//	}
//}

//type APIResponse struct {
//	Success bool        `json:"success"`
//	Data    interface{} `json:"data,omitempty"`
//	Error   string      `json:"error,omitempty"`
//	Message string      `json:"message,omitempty"`
//}

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

func (h *Handler) AddBook(w http.ResponseWriter, r *http.Request) {

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

func (h *Handler) ListOrSearchBooks(w http.ResponseWriter, r *http.Request, params api.ListOrSearchBooksParams) {
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

	allBooks, err := h.bookService.GetAllBooks(r.Context(), paginationParams)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	books := make([]api.Books, len(allBooks.Results))
	for i, book := range allBooks.Results {
		books[i] = *book
	}

	apiPagination := &api.PaginationInfo{
		Offset:      &allBooks.Pagination.Offset,
		Limit:       &allBooks.Pagination.Limit,
		Total:       &allBooks.Pagination.Total,
		HasNext:     &allBooks.Pagination.HasNext,
		HasPrevious: &allBooks.Pagination.HasPrevious,
	}

	h.writeResponse(w, http.StatusOK, api.ListOrSearchBooks200JSONResponse{Results: &books, Pagination: apiPagination})

}

func (h *Handler) ListOverdueRentals(w http.ResponseWriter, r *http.Request, params api.ListOverdueRentalsParams) {
	panic("implement me")
}

func (h *Handler) GetRentalReports(w http.ResponseWriter, r *http.Request, params api.GetRentalReportsParams) {
	panic("implement me")
}
