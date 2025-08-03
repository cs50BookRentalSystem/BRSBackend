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

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) repository.CartRepository {
	return &cartRepository{db: db}
}

func (c cartRepository) Create(ctx context.Context, cart *models.Cart) error {
	if err := c.db.WithContext(ctx).Create(cart).Error; err != nil {
		return fmt.Errorf("failed to create cart: %w", err)
	}

	return nil
}

func (c cartRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	if err := c.db.WithContext(ctx).Where("id = ?", id).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("cart not found")
		}
		return nil, fmt.Errorf("failed to find cart: %w", err)
	}

	return &cart, nil
}

func (c cartRepository) GetByStatus(ctx context.Context, status string) ([]*models.Cart, error) {
	var carts []*models.Cart
	if err := c.db.WithContext(ctx).Where("status = ?", status).Find(&carts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("cart not found")
		}
	}
	return carts, nil
}

func (c cartRepository) GetCartsByStudentID(ctx context.Context, studentID uuid.UUID) ([]*models.Cart, error) {
	var carts []*models.Cart
	return carts, c.db.WithContext(ctx).Where("student_id = ?", studentID).Find(&carts).Error
}

func (c cartRepository) Update(ctx context.Context, cart *models.Cart) error {
	if err := c.db.WithContext(ctx).Save(cart).Error; err != nil {
		return fmt.Errorf("failed to update cart: %w", err)
	}
	return nil
}

func (c cartRepository) UpdateStatus(ctx context.Context, cartID uuid.UUID, status string) error {
	if err := c.db.WithContext(ctx).Model(&models.Cart{}).
		Where("id = ?", cartID).
		Update("status", status).Error; err != nil {
		return fmt.Errorf("failed to update cart status: %w", err)
	}
	return nil
}
