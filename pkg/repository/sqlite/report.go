package sqlite

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) repository.ReportRepository {
	return &reportRepository{db: db}
}

func (r reportRepository) GetOverdueRentals(ctx context.Context, studentCardID *string, limit, offset int) ([]dto.OverdueUser, int64, error) {
	var overdueUsers []dto.OverdueUser
	var total int64

	overduePeriod := 7 * 24 * time.Hour
	overdueDate := time.Now().Add(-overduePeriod)

	query := r.db.WithContext(ctx).
		Table("rents").
		Select(`
			students.first_name || ' ' || students.last_name as student_name,
			students.phone,
			COUNT(rents.id) as total_books,
			MIN(carts.created_at) as date_rented,
			CAST((julianday('now') - julianday(MIN(carts.created_at))) as INTEGER) as days_overdue
		`).
		Joins("JOIN carts ON rents.cart_id = carts.id").
		Joins("JOIN students ON carts.student_id = students.id").
		Where("carts.status = ? AND carts.created_at < ?", "RENTED", overdueDate).
		Group("students.id, students.first_name, students.last_name, students.phone").
		Having("days_overdue > ?", 7)

	if studentCardID != nil && *studentCardID != "" {
		query = query.Where("students.card_id = ?", *studentCardID)
	}

	countQuery := query
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count overdue rentals: %w", err)
	}

	if err := query.Limit(limit).Offset(offset).Scan(&overdueUsers).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get overdue rentals: %w", err)
	}

	return overdueUsers, total, nil
}

func (r reportRepository) GetRentalReport(ctx context.Context, limit, offset int) (*dto.RentReport, error) {
	var report dto.RentReport
	var totalRents int64
	var totalStudents int64
	var topBooks []dto.BookRentStats

	if err := r.db.WithContext(ctx).Model(&models.Rent{}).Count(&totalRents).Error; err != nil {
		return nil, fmt.Errorf("failed to count total rents: %w", err)
	}
	report.TotalRents = int(totalRents)

	if err := r.db.WithContext(ctx).
		Table("students").
		Joins("JOIN carts ON students.id = carts.student_id").
		Distinct("students.id").
		Count(&totalStudents).Error; err != nil {
		return nil, fmt.Errorf("failed to count total students: %w", err)
	}
	report.TotalStudents = int(totalStudents)

	if err := r.db.WithContext(ctx).
		Table("rents").
		Select("books.title as book_title, COUNT(rents.id) as rented_count").
		Joins("JOIN books ON rents.book_id = books.id").
		Group("books.id, books.title").
		Order("rented_count DESC").
		Limit(limit).
		Offset(offset).
		Scan(&topBooks).Error; err != nil {
		return nil, fmt.Errorf("failed to get top books: %w", err)
	}
	report.TopBooks = topBooks

	return &report, nil
}
