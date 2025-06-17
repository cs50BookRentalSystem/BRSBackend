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

func (r rentRepository) GetRentsByFilters(ctx context.Context, filters dto.RentFilters) ([]*dto.RentSummary, error) {
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

	if filters.Limit > 0 {
		query = query.Limit(filters.Limit)
	}
	if filters.Offset > 0 {
		query = query.Offset(filters.Offset)
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to get rents by filters: %w", err)
	}

	return results, nil
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

func (r rentRepository) GetOverdueRentals(ctx context.Context, studentCardID string, offset, limit int) ([]*dto.OverdueUser, int64, error) {
	//var overdueUsers []*repository.OverdueUser
	//var total int64
	//
	//// Assuming books are overdue after 14 days
	//overdueDate := time.Now().AddDate(0, 0, -14)
	//
	//baseQuery := `
	//	FROM carts
	//	JOIN students ON carts.student_id = students.id
	//	JOIN rents ON carts.id = rents.cart_id
	//	WHERE carts.status = 'RENTED' AND carts.created_at < ?
	//`
	//
	//args := []interface{}{overdueDate}
	//if studentCardID != "" {
	//	baseQuery += " AND students.card_id = ?"
	//	args = append(args, studentCardID)
	//}
	//
	//countQuery := "SELECT COUNT(DISTINCT students.id) " + baseQuery
	//if err := r.db.WithContext(ctx).Raw(countQuery, args...).Scan(&total).Error; err != nil {
	//	return nil, 0, fmt.Errorf("failed to count overdue users: %w", err)
	//}
	//
	//query := `
	//	SELECT
	//		(students.first_name || ' ' || students.last_name) as student_name,
	//		students.phone,
	//		COUNT(rents.id) as total_books,
	//		MIN(carts.created_at) as date_rented,
	//		CAST(julianday('now') - julianday(MIN(carts.created_at)) as INTEGER) as days_overdue
	//	` + baseQuery + `
	//	GROUP BY students.id, students.first_name, students.last_name, students.phone
	//	ORDER BY days_overdue DESC
	//	LIMIT ? OFFSET ?
	//`
	//
	//args = append(args, limit, offset)
	//if err := r.db.WithContext(ctx).Raw(query, args...).Scan(&overdueUsers).Error; err != nil {
	//	return nil, 0, fmt.Errorf("failed to get overdue users: %w", err)
	//}
	//
	//return overdueUsers, total, nil
	panic("implement me")
}

func (r rentRepository) GetRentalReport(ctx context.Context) (*dto.RentReport, error) {
	//var report repository.RentReport
	//
	//if err := r.db.WithContext(ctx).Model(&models.Rent{}).Count(&report.TotalRents).Error; err != nil {
	//	return nil, fmt.Errorf("failed to count total rents: %w", err)
	//}
	//
	//if err := r.db.WithContext(ctx).Model(&models.Student{}).Count(&report.TotalStudents).Error; err != nil {
	//	return nil, fmt.Errorf("failed to count total students: %w", err)
	//}
	//
	//var topBooks []*repository.BookRentStats
	//query := `
	//	SELECT
	//		books.title as book_title,
	//		COUNT(rents.id) as rented_count
	//	FROM books
	//	JOIN rents ON books.id = rents.book_id
	//	GROUP BY books.id, books.title
	//	ORDER BY rented_count DESC
	//	LIMIT 10
	//`
	//
	//if err := r.db.WithContext(ctx).Raw(query).Scan(&topBooks).Error; err != nil {
	//	return nil, fmt.Errorf("failed to get top books: %w", err)
	//}
	//
	//report.TopBooks = topBooks
	//return &report, nil
	panic("implement me")
}
