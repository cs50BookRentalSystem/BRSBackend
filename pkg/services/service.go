package services

import (
	"BRSBackend/pkg/repository"
)

type Service struct {
	Book BookService
	Auth AuthService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repo.Book),
		Auth: NewAuthService(repo.Librarian, repo.Session),
	}
}
