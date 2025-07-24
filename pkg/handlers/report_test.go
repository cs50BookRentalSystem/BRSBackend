package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/services"
)

func TestGetRentalReports(t *testing.T) {
	t.Run("successful get rental reports", func(t *testing.T) {
		mockReportService := &services.MockReportService{
			GetRentalReportFunc: func(ctx context.Context, limit, offset int) (*dto.RentReport, error) {
				return &dto.RentReport{
					TotalRents:    100,
					TotalStudents: 50,
					TopBooks: []dto.BookRentStats{
						{BookTitle: "Test Book 1", RentedCount: 10},
						{BookTitle: "Test Book 2", RentedCount: 5},
					},
					TopOverdue: []dto.OverdueUser{},
				}, nil
			},
		}

		h := NewHandler(&services.Service{Report: mockReportService})

		req := httptest.NewRequest(http.MethodGet, "/reports", nil)
		w := httptest.NewRecorder()

		h.GetRentalReports(w, req, api.GetRentalReportsParams{})

		if w.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
		}

		var report dto.RentReport
		if err := json.NewDecoder(w.Body).Decode(&report); err != nil {
			t.Errorf("failed to decode response body: %v", err)
		}

		if report.TotalRents != 100 {
			t.Errorf("expected total rents to be 100, got %d", report.TotalRents)
		}

		if report.TotalStudents != 50 {
			t.Errorf("expected total students to be 50, got %d", report.TotalStudents)
		}

		if len(report.TopBooks) != 2 {
			t.Errorf("expected 2 top books, got %d", len(report.TopBooks))
		}
	})

	t.Run("successful get rental reports with overdue data", func(t *testing.T) {
		mockReportService := &services.MockReportService{
			GetRentalReportFunc: func(ctx context.Context, limit, offset int) (*dto.RentReport, error) {
				return &dto.RentReport{
					TotalRents:    100,
					TotalStudents: 50,
					TopBooks:      []dto.BookRentStats{},
					TopOverdue: []dto.OverdueUser{
						{
							StudentName: "John Doe",
							Phone:       "1234567890",
							TotalBooks:  1,
							DateRented:  time.Now(),
							DaysOverdue: 10,
						},
					},
				}, nil
			},
		}

		h := NewHandler(&services.Service{Report: mockReportService})

		req := httptest.NewRequest(http.MethodGet, "/reports", nil)
		w := httptest.NewRecorder()

		h.GetRentalReports(w, req, api.GetRentalReportsParams{})

		if w.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
		}

		var report dto.RentReport
		if err := json.NewDecoder(w.Body).Decode(&report); err != nil {
			t.Errorf("failed to decode response body: %v", err)
		}

		if len(report.TopOverdue) != 1 {
			t.Errorf("expected 1 top overdue user, got %d", len(report.TopOverdue))
		}

		if report.TopOverdue[0].StudentName != "John Doe" {
			t.Errorf("expected overdue student name to be 'John Doe', got '%s'", report.TopOverdue[0].StudentName)
		}
	})
}
