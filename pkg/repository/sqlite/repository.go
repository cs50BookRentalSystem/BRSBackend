package sqlite

import (
	"gorm.io/gorm"

	"BRSBackend/pkg/repository"
)

func NewRepository(db *gorm.DB) *repository.Repository {
	return &repository.Repository{
		Book:      NewBookRepository(db),
		Student:   NewStudentRepository(db),
		Librarian: NewLibrarianRepository(db),
		Cart:      NewCartRepository(db),
		Rent:      NewRentRepository(db),
		Session:   NewSessionRepository(db),
		Report:    NewReportRepository(db),
	}
}
