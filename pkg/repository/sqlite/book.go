package sqlite

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type bookRepository struct {
	db *gorm.DB
}

func (b bookRepository) Create(ctx context.Context, book *models.Book) error {
	if err := b.db.WithContext(ctx).Create(book).Error; err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}

	return nil
}

func (b bookRepository) GetByID(ctx context.Context, id string) (*models.Book, error) {
	var book models.Book
	if err := b.db.WithContext(ctx).Where("id = ?", id).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found")
		}
		return nil, fmt.Errorf("failed to get book: %w", err)
	}

	return &book, nil
}

func (b bookRepository) GetAll(ctx context.Context, offset, limit int) ([]*models.Book, int64, error) {
	var books []*models.Book
	var total int64

	if err := b.db.WithContext(ctx).Model(&models.Book{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get book count: %w", err)
	}

	if err := b.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&books).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get books: %w", err)
	}

	return books, total, nil
}

func (b bookRepository) SearchByTitle(ctx context.Context, title string, offset, limit int) ([]*models.Book, int64, error) {
	var books []*models.Book
	var total int64

	query := b.db.WithContext(ctx).Model(&models.Book{}).
		Where("title LIKE ? OR id = ?", "%"+title+"%", title)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get book count: %w", err)
	}

	if err := query.Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&books).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get books: %w", err)
	}

	return books, total, nil
}

func (b bookRepository) Update(ctx context.Context, book *models.Book) error {
	panic("implement me")
}

func (b bookRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{db: db}
}
