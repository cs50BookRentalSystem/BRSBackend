package services

import (
	"context"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
)

type StudentService interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	GetStudentByID(ctx context.Context, id string) (*models.Student, error)
	GetAllStudents(ctx context.Context, params dto.PaginationParams) (*dto.StudentsResponse, error)
}

type PaginatedStudentsResponse struct {
	Books      []*models.Student `json:"students"`
	Page       int               `json:"page"`
	PageSize   int               `json:"page_size"`
	TotalItems int               `json:"total_items"`
	HasMore    bool              `json:"has_more"`
}
