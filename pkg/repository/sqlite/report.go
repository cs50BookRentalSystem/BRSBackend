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

func (r reportRepository) GetOverdueRentals(ctx context.Context, studentCardID *string, limit, offset, overduePeriod int) ([]dto.OverdueUser, int64, error) {
	var overdueUsers []dto.OverdueUser
	var total int64

	query := r.db.WithContext(ctx).
		Table("rents").
		Select(`
			students.id,
			carts.id as cart_id,
			students.first_name || ' ' || students.last_name as student_name,
	 		students.card_id,
			students.phone,
			COUNT(*) as total_books,
			MIN(carts.created_at) as date_rented,
			julianday('now') - julianday(MIN(carts.created_at)) as days_overdue
		`).
		Joins("JOIN carts ON rents.cart_id = carts.id").
		Joins("JOIN students ON carts.student_id = students.id").
		Where("carts.status = ? ", "RENTED").
		Group("students.id, students.first_name, students.last_name, students.phone").
		Having("julianday('now') - julianday(MIN(carts.created_at)) > ?", overduePeriod)

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
		CartId      string  `json:"cart_id"`
		StudentName string  `json:"student_name"`
		CardId      string  `json:"card_id"`
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
		dateRented, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", temp.DateRented)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to parse date_rented '%s': %w", temp.DateRented, err)
		}

		daysOverdue := int(temp.DaysOverdue)

		if daysOverdue > overduePeriod {
			daysOverdue = daysOverdue - overduePeriod
		}

		overdueUsers[i] = dto.OverdueUser{
			CartId:      temp.CartId,
			StudentName: temp.StudentName,
			CardId:      temp.CardId,
			Phone:       temp.Phone,
			TotalBooks:  temp.TotalBooks,
			DateRented:  dateRented,
			DaysOverdue: daysOverdue,
		}
	}

	return overdueUsers, total, nil
}

func (r reportRepository) GetRentalReport(ctx context.Context, limit, offset, overduePeriod int) (*dto.RentReport, error) {
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

	overdueUsers, _, err := r.GetOverdueRentals(ctx, nil, limit, offset, overduePeriod)
	if err != nil {
		return nil, fmt.Errorf("failed to get overdue rentals: %w", err)
	}
	report.TopOverdue = overdueUsers

	return &report, nil
}
