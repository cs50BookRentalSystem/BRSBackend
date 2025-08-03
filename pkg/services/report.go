package services

import (
	"context"
	"fmt"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/repository"
)

type ReportService interface {
	GetOverdueRentals(ctx context.Context, studentCardID *string, limit, offset int) (*dto.OverdueResponse, error)
	GetRentalReport(ctx context.Context, limit, offset int) (*dto.RentReport, error)
}

type reportService struct {
	repo          repository.ReportRepository
	overduePeriod int
}

func NewReportService(repo repository.ReportRepository, overduePeriod int) ReportService {
	return &reportService{
		repo:          repo,
		overduePeriod: overduePeriod,
	}
}

func (r *reportService) GetOverdueRentals(ctx context.Context, studentCardID *string, limit, offset int) (*dto.OverdueResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	overdueUsers, total, err := r.repo.GetOverdueRentals(ctx, studentCardID, limit, offset, r.overduePeriod)
	if err != nil {
		return nil, fmt.Errorf("failed to get overdue rentals: %w", err)
	}

	return &dto.OverdueResponse{
		Results: overdueUsers,
		Pagination: dto.PaginationInfo{
			Offset:      offset,
			Limit:       limit,
			Total:       int(total),
			HasNext:     offset+limit < len(overdueUsers),
			HasPrevious: offset > 0,
		},
	}, nil

}

func (r *reportService) GetRentalReport(ctx context.Context, limit, offset int) (*dto.RentReport, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	report, err := r.repo.GetRentalReport(ctx, limit, offset, r.overduePeriod)
	if err != nil {
		return nil, fmt.Errorf("failed to get rental report: %w", err)
	}

	return report, nil
}
