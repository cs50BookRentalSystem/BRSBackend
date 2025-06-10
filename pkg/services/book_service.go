package services

import (
	"context"

	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, book *models.Book) (*models.Book, error)
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	GetAllBooks(ctx context.Context, limit, offset int) ([]*models.Book, error)
	SearchByTitle(ctx context.Context, title string, offset, limit int) ([]*models.Book, int64, error)
}

type bookService struct {
	repo repository.BookRepository
}

func (b bookService) CreateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) GetAllBooks(ctx context.Context, limit, offset int) ([]*models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookService) SearchByTitle(ctx context.Context, title string, offset, limit int) ([]*models.Book, int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}
