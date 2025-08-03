package services

import (
	"BRSBackend/pkg/repository"
)

type Service struct {
	Book    BookService
	Auth    AuthService
	Student StudentService
	Rent    RentService
	Report  ReportService
}

func NewService(repo *repository.Repository, overduePeriod int) *Service {
	return &Service{
		Book:    NewBookService(repo.Book),
		Auth:    NewAuthService(repo.Librarian, repo.Session),
		Student: NewStudentService(repo.Student),
		Rent:    NewRentService(repo.Rent, repo.Cart, repo.Book, repo.Student),
		Report:  NewReportService(repo.Report, overduePeriod),
	}
}
