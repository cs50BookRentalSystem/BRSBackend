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

	//overduePeriod := 1 * 24 * time.Hour
	//overdueDate := time.Now().Add(-overduePeriod)

	query := r.db.WithContext(ctx).
		Table("rents").
		Select(`
			students.id,
			students.first_name || ' ' || students.last_name as student_name,
			students.phone,
			COUNT(*) as total_books,
			MIN(carts.created_at) as date_rented,
			julianday('now') - julianday(MIN(carts.created_at)) as days_overdue
		`).
		Joins("JOIN carts ON rents.cart_id = carts.id").
		Joins("JOIN students ON carts.student_id = students.id").
		Where("carts.status = ? ", "RENTED").
		Group("students.id, students.first_name, students.last_name, students.phone").
		Having("julianday('now') - julianday(MIN(carts.created_at)) > ?", 1)

	if studentCardID != nil && *studentCardID != "" {
		query = query.Where("students.card_id = ?", *studentCardID)
	}

	var countResult struct {
		Count int64
	}

	countQuery := r.db.WithContext(ctx).
		Table("(?) as grouped_results", query).
		Select("COUNT(*) as count")

	type tempOverdueUser struct {
		ID          string  `json:"id"`
		StudentName string  `json:"student_name"`
		Phone       string  `json:"phone"`
		TotalBooks  int     `json:"total_books"`
		DateRented  string  `json:"date_rented"`
		DaysOverdue float64 `json:"days_overdue"`
	}
	var tempUsers []tempOverdueUser

	if err := countQuery.Scan(&countResult).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count overdue rentals: %w", err)
	}
	total = countResult.Count

	if err := query.Limit(limit).Offset(offset).Scan(&tempUsers).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get overdue rentals: %w", err)
	}
	overdueUsers = make([]dto.OverdueUser, len(tempUsers))
	for i, temp := range tempUsers {
		dateRented, err := time.Parse("2006-01-02 15:04:05.000000-07:00", temp.DateRented)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to parse date_rented '%s': %w", temp.DateRented, err)
		}

		overdueUsers[i] = dto.OverdueUser{
			StudentName: temp.StudentName,
			Phone:       temp.Phone,
			TotalBooks:  temp.TotalBooks,
			DateRented:  dateRented,
			DaysOverdue: int(temp.DaysOverdue),
		}
	}

	return overdueUsers, total, nil
}

//SELECT count(*) FROM `rents` JOIN carts ON rents.cart_id = carts.id JOIN students ON carts.student_id = students.id WHERE (carts.status = "RENTED" AND carts.created_at < "2025-06-17 11:15:29.603") AND students.card_id = "1234567890" GROUP BY students.id, students.first_name, students.last_name, students.phone HAVING days_overdue > 1;
//SELECT
//students.id,
//students.first_name || ' ' || students.last_name as student_name,
//students.phone,
//COUNT(*) as total_books,
//MIN(carts.created_at) as date_rented,
//julianday('now') - julianday(MIN(carts.created_at)) as days_overdue
//FROM `rents` JOIN carts ON rents.cart_id = carts.id JOIN students ON carts.student_id = students.id WHERE (carts.status = "RENTED" AND carts.created_at < "2025-06-18 19:54:46.039") AND students.card_id = "1234567890" GROUP BY students.id, students.first_name, students.last_name, students.phone HAVING julianday('now') - julianday(MIN(carts.created_at)) > 1 LIMIT 20

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
		Group("books.title").
		Order("rented_count DESC").
		Limit(limit).
		Offset(offset).
		Scan(&topBooks).Error; err != nil {
		return nil, fmt.Errorf("failed to get top books: %w", err)
	}
	report.TopBooks = topBooks

	overdueUsers, _, err := r.GetOverdueRentals(ctx, nil, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get overdue rentals: %w", err)
	}
	report.TopOverdue = overdueUsers

	return &report, nil
}

// SELECT books.title as book_title, COUNT(rents.id) as rented_count FROM `rents` JOIN books ON rents.book_id = books.id GROUP BY books.id, books.title ORDER BY rented_count DESC LIMIT 20
