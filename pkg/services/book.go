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

type CreateBookRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required,min=1"`
	Count       int    `json:"count" validate:"required,min=0"`
}

type UpdateBookRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=1"`
	Count       *int    `json:"count,omitempty" validate:"omitempty,min=0"`
}

type PaginatedBooksResponse struct {
	Books      []*models.Book `json:"books"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalItems int            `json:"total_items"`
	HasMore    bool           `json:"has_more"`
}

type bookService struct {
	repo repository.BookRepository
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

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}
