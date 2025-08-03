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

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{db: db}
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
			search := fmt.Sprintf("%%%s%%", strings.ToLower(strings.Trim(params.Query, `"`)))
			query = query.Where("LOWER(title) LIKE ?", search)
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

func (b *bookRepository) GetBooksByIDs(ctx context.Context, bookIDs []uuid.UUID) ([]*models.Book, error) {
	var books []*models.Book
	return books, b.db.WithContext(ctx).Where("id IN ?", bookIDs).Find(&books).Error
}

func (b *bookRepository) UpdateCount(ctx context.Context, bookID uuid.UUID, delta int) error {

	tx := b.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var book models.Book
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("id = ?", bookID).First(&book).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book not found")
		}
		return fmt.Errorf("failed to get book: %w", err)
	}

	newCount := book.Count + delta
	if newCount < 0 {
		tx.Rollback()
		return fmt.Errorf("insufficient book count: current=%d, requested=%d", book.Count, -delta)
	}

	if err := tx.Model(&book).Update("count", newCount).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update book count: %w", err)
	}

	return tx.Commit().Error
}

func (b *bookRepository) DecrementCount(ctx context.Context, bookID uuid.UUID) error {
	return b.UpdateCount(ctx, bookID, -1)
}

func (b *bookRepository) IncrementCount(ctx context.Context, bookID uuid.UUID) error {
	return b.UpdateCount(ctx, bookID, 1)
}

func (b *bookRepository) DecrementMultipleBooks(ctx context.Context, bookIDs []uuid.UUID) error {
	tx := b.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	bookCounts := make(map[uuid.UUID]int)
	for _, bookID := range bookIDs {
		bookCounts[bookID]++
	}

	for bookID, count := range bookCounts {
		var book models.Book
		if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("id = ?", bookID).First(&book).Error; err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("book not found: %s", bookID)
			}
			return fmt.Errorf("failed to get book %s: %w", bookID, err)
		}

		newCount := book.Count - count
		if newCount < 0 {
			tx.Rollback()
			return fmt.Errorf("insufficient book count for '%s': current=%d, requested=%d", book.Title, book.Count, count)
		}

		if err := tx.Model(&book).Update("count", newCount).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update book count for %s: %w", bookID, err)
		}
	}

	return tx.Commit().Error
}

func (b *bookRepository) IncrementMultipleBooks(ctx context.Context, bookIDs []uuid.UUID) error {

	tx := b.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	bookCounts := make(map[uuid.UUID]int)
	for _, bookID := range bookIDs {
		bookCounts[bookID]++
	}

	for bookID, count := range bookCounts {
		if err := tx.Model(&models.Book{}).Where("id = ?", bookID).
			Update("count", gorm.Expr("count + ?", count)).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to increment book count for %s: %w", bookID, err)
		}
	}

	return tx.Commit().Error
}

func (b *bookRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return b.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Book{}).Error
}
