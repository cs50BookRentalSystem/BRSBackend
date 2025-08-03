package dto

import "time"

type OverdueUser struct {
	CartId      string    `json:"cartId"`
	StudentName string    `json:"student_name"`
	CardId      string    `json:"card_id"`
	Phone       string    `json:"phone"`
	TotalBooks  int       `json:"total_books"`
	DateRented  time.Time `json:"date_rented"`
	DaysOverdue int       `json:"days_overdue"`
}

type BookRentStats struct {
	BookTitle   string `json:"book_title"`
	RentedCount int    `json:"rented_count"`
}

type RentReport struct {
	TotalRents    int             `json:"total_rents"`
	TotalStudents int             `json:"total_students"`
	TopBooks      []BookRentStats `json:"top_books"`
	TopOverdue    []OverdueUser   `json:"top_overdue"`
}

type OverdueResponse struct {
	Results    []OverdueUser  `json:"results"`
	Pagination PaginationInfo `json:"pagination"`
}
