package dto

import "BRSBackend/pkg/models"

type PaginationParams struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Query  string `json:"query,omitempty"`
}

type PaginationInfo struct {
	Offset      int  `json:"offset"`
	Limit       int  `json:"limit"`
	Total       int  `json:"total"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}

type BooksResponse struct {
	Results    []*models.Book `json:"results"`
	Pagination PaginationInfo `json:"pagination"`
}

type StudentsResponse struct {
	Results    []*models.Student `json:"results"`
	Pagination PaginationInfo    `json:"pagination"`
}
