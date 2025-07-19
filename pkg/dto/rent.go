package dto

import (
	"time"

	"github.com/google/uuid"
)

type RentSummary struct {
	RentID      uuid.UUID `json:"rent_id"`
	CartID      uuid.UUID `json:"cart_id"`
	BookTitle   string    `json:"book_title"`
	StudentName string    `json:"student_name"`
	RentedDate  time.Time `json:"rented_date"`
}

type RentFilters struct {
	BookName    *string
	StudentName *string
	Date        *time.Time
	Limit       int
	Offset      int
}

type CreateRentRequest struct {
	StudentID uuid.UUID   `json:"student_id" validate:"required"`
	BookIDs   []uuid.UUID `json:"book_ids" validate:"required,min=1"`
}

type CreateRentResponse struct {
	CartID  uuid.UUID `json:"cart_id"`
	Message string    `json:"message"`
}

type GetRentedBooksResponse struct {
	Results    []*RentSummary `json:"results"`
	Pagination PaginationInfo `json:"pagination"`
}

type ReturnBooksResponse struct {
	Message string    `json:"message"`
	CartID  uuid.UUID `json:"cart_id"`
}
