package sqlite

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type bookRepository struct {
	db *gorm.DB
}

func (b *bookRepository) Create(ctx context.Context, book *models.Book) error {
	if err := b.db.WithContext(ctx).Create(book).Error; err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}

	return nil
}

func (b *bookRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Book, error) {
	var book models.Book
	if err := b.db.WithContext(ctx).Where("id = ?", id).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, fmt.Errorf("failed to get book: %w", err)
	}

	return &book, nil
}

func (b *bookRepository) GetAll(ctx context.Context, params dto.PaginationParams) ([]*models.Book, int64, error) {
	var books []*models.Book
	var total int64

	query := b.db.WithContext(ctx).Model(&models.Book{})

	if params.Query != "" {
		if id, err := uuid.Parse(params.Query); err == nil {
			query = query.Where("id = ?", id)
		} else {
			query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(params.Query)+"%")
		}
	}

	if err := query.Count(&total).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, fmt.Errorf("book not found")
	}

	if err := query.
		Order("created_at DESC").
		Offset(params.Offset).
		Limit(params.Limit).
		Find(&books).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get books: %w", err)
	}

	return books, total, nil
}

func (b *bookRepository) Update(ctx context.Context, book *models.Book) error {
	panic("implement me")
}

func (b *bookRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{db: db}
}
