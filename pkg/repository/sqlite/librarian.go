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

type librarianRepository struct {
	db *gorm.DB
}

func NewLibrarianRepository(db *gorm.DB) repository.LibrarianRepository {
	return &librarianRepository{db: db}
}

func (l *librarianRepository) Create(ctx context.Context, librarian *models.Librarian) error {
	if err := l.db.WithContext(ctx).Create(librarian).Error; err != nil {
		return fmt.Errorf("failed to create librarian: %w", err)
	}
	return nil
}

func (l *librarianRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Librarian, error) {
	var librarian models.Librarian
	if err := l.db.WithContext(ctx).Where("id = ?", id).First(&librarian).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("librarian not found")
		}
		return nil, fmt.Errorf("failed to get librarian by id: %w", err)
	}
	return &librarian, nil
}

func (l *librarianRepository) GetByUsername(ctx context.Context, username string) (*models.Librarian, error) {
	var librarian models.Librarian
	if err := l.db.WithContext(ctx).Where("user = ?", username).First(&librarian).Error; err != nil {
		return nil, fmt.Errorf("failed to get librarian by username: %w", err)
	}
	return &librarian, nil
}
