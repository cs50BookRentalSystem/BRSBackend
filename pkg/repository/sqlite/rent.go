package sqlite

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type rentRepository struct {
	db *gorm.DB
}

func NewRentRepository(db *gorm.DB) repository.RentRepository {
	return &rentRepository{db: db}
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

func (r rentRepository) GetRentsByFilters(ctx context.Context, filters dto.RentFilters) ([]*dto.RentSummary, int64, error) {
	var results []*dto.RentSummary
	var total int64

	query := r.db.WithContext(ctx).
		Table("rents").
		Select(`
			rents.id as rent_id,
			rents.cart_id,
			books.title as book_title,
			CONCAT(students.first_name, ' ', students.last_name) as student_name,
			carts.created_at as rented_date
		`).
		Joins("JOIN carts ON rents.cart_id = carts.id").
		Joins("JOIN books ON rents.book_id = books.id").
		Joins("JOIN students ON carts.student_id = students.id").
		Where("carts.status = ?", "RENTED")

	if filters.BookName != nil && *filters.BookName != "" {
		searchBook := fmt.Sprintf("%%%s%%", strings.ToLower(strings.Trim(*filters.BookName, `"`)))
		query = query.Where("books.title LIKE ?", searchBook)
	}

	if filters.StudentName != nil && *filters.StudentName != "" {
		fullName := strings.ToLower(strings.Trim(*filters.StudentName, `"`))
		if fullName != "" {
			parts := strings.Fields(fullName)

			switch len(parts) {
			case 1:
				searchTerm := "%" + parts[0] + "%"
				query = query.Where(
					"LOWER(students.first_name) LIKE ? OR LOWER(students.last_name) LIKE ?",
					searchTerm, searchTerm,
				)
			default:
				firstName := parts[0]
				lastName := parts[len(parts)-1]

				query = query.Where(
					"LOWER(students.first_name) LIKE ? OR LOWER(students.last_name) LIKE ? OR LOWER(CONCAT(students.first_name, ' ', students.last_name)) LIKE ?",
					"%"+firstName+"%", "%"+lastName+"%", "%"+fullName+"%",
				)
			}
		}
	}

	if filters.Date != nil {
		startOfDay := filters.Date.Truncate(24 * time.Hour)
		endOfDay := startOfDay.Add(24 * time.Hour)
		query = query.Where("carts.created_at >= ? AND carts.created_at < ?", startOfDay, endOfDay)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get total rents by filters: %w", err)
	}

	if filters.Limit > 0 {
		query = query.Limit(filters.Limit)
	}
	if filters.Offset > 0 {
		query = query.Offset(filters.Offset)
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get rents by filters: %w", err)
	}

	return results, total, nil
}

func (r rentRepository) GetRentedBooksByStudent(ctx context.Context, studentCardID string) ([]*dto.RentSummary, error) {
	var results []*dto.RentSummary

	query := r.db.WithContext(ctx).
		Table("rents").
		Select(`
			rents.id as rent_id,
			rents.cart_id,
			books.title as book_title,
			CONCAT(students.first_name, ' ', students.last_name) as student_name,
			carts.created_at as rented_date
		`).
		Joins("JOIN carts ON rents.cart_id = carts.id").
		Joins("JOIN books ON rents.book_id = books.id").
		Joins("JOIN students ON carts.student_id = students.id").
		Where("carts.status = ?", "RENTED")

	if studentCardID != "" {
		query = query.Where("students.card_id = ?", studentCardID)
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to get rented books by student card: %w", err)
	}

	return results, nil
}

func (r rentRepository) GetRentsByCartID(ctx context.Context, cartID uuid.UUID) ([]*models.Rent, error) {
	var rents []*models.Rent

	if err := r.db.WithContext(ctx).Where("cart_id = ?", cartID).Find(&rents).Error; err != nil {
		return nil, fmt.Errorf("failed to get rents by cart ID: %w", err)
	}

	return rents, nil
}
