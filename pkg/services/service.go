package services

import (
	"BRSBackend/pkg/repository"
)

type Service struct {
	Book    BookService
	Auth    AuthService
	Student StudentService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Book:    NewBookService(repo.Book),
		Auth:    NewAuthService(repo.Librarian, repo.Session),
		Student: NewStudentService(repo.Student),
	}
}
