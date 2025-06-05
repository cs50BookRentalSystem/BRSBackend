package sqlite

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type rentRepository struct {
	db *gorm.DB
}

func (r rentRepository) Create(ctx context.Context, rent *models.Rent) error {
	if err := r.db.WithContext(ctx).Create(rent).Error; err != nil {
		return fmt.Errorf("failed to create rent: %w", err)
	}

	return nil
}

func (r rentRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Rent, error) {
	var rent models.Rent
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&rent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("rent not found")
		}
		return nil, fmt.Errorf("failed to get rent: %w", err)
	}

	return &rent, nil
}

func (r rentRepository) GetByCartId(ctx context.Context, cartID uuid.UUID) ([]*models.Rent, error) {
	var rents []*models.Rent
	if err := r.db.WithContext(ctx).Where("cart_id = ?", cartID).Find(&rents).Error; err != nil {
		return nil, fmt.Errorf("failed to get rents by cart ID: %w", err)
	}

	return rents, nil
}

func (r rentRepository) GetAll(ctx context.Context, offset, limit int) ([]*models.Rent, int64, error) {
	var rents []*models.Rent
	var total int64

	if err := r.db.WithContext(ctx).Model(&models.Rent{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count rents: %w", err)
	}
	if err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&rents).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get rents: %w", err)
	}

	return rents, total, nil
}

func (r rentRepository) GetRentedBooks(ctx context.Context, studentCardID string) ([]*repository.RentSummary, error) {
	//TODO implement me
	panic("implement me")
}

func (r rentRepository) GetOverdueRentals(ctx context.Context, studentCardID string, offset, limit int) ([]*repository.OverdueUser, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r rentRepository) GetRentalReport(ctx context.Context) (*repository.RentReport, error) {
	//TODO implement me
	panic("implement me")
}

func (r rentRepository) SearchRents(ctx context.Context, bookName, studentName string, date *time.Time, offset, limit int) ([]*repository.RentSummary, int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewRentRepository(db *gorm.DB) repository.RentRepository {
	return &rentRepository{db: db}
}
