package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, book *models.Book) error
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	GetAllBooks(ctx context.Context, params dto.PaginationParams) (*dto.BooksResponse, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

func (b *bookService) CreateBook(ctx context.Context, book *models.Book) error {

	if book.Title == "" {
		return fmt.Errorf("book title is required")
	}
	if book.Count < 0 {
		return fmt.Errorf("book cannot be negative")
	}

	return b.repo.Create(ctx, book)
}

func (b *bookService) GetBookByID(ctx context.Context, uid string) (*models.Book, error) {

	id, err := uuid.Parse(uid)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %w", err)
	}

	return b.repo.GetByID(ctx, id)
}

func (b *bookService) GetAllBooks(ctx context.Context, params dto.PaginationParams) (*dto.BooksResponse, error) {

	if params.Limit <= 0 {
		params.Limit = 10
	}
	if params.Limit > 100 {
		params.Limit = 100
	}
	if params.Offset < 0 {
		params.Offset = 0
	}

	books, total, err := b.repo.GetAll(ctx, params)
	if err != nil {
		return nil, err
	}

	pagination := dto.PaginationInfo{
		Offset:      params.Offset,
		Limit:       params.Limit,
		Total:       int(total),
		HasNext:     int64(params.Offset+params.Limit) < total,
		HasPrevious: params.Offset > 0,
	}

	response := &dto.BooksResponse{
		Results:    books,
		Pagination: pagination,
	}

	return response, nil
}
