package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type StudentService interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	GetStudentByID(ctx context.Context, id string) (*models.Student, error)
	GetStudentByCardNumber(ctx context.Context, number string) (*models.Student, error)
	GetAllStudents(ctx context.Context, params dto.PaginationParams) (*dto.StudentsResponse, error)
	DeleteStudent(ctx context.Context, id string) error
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{
		repo: repo,
	}
}

func (s *studentService) CreateStudent(ctx context.Context, student *models.Student) error {
	return s.repo.Create(ctx, student)
}

func (s *studentService) GetStudentByID(ctx context.Context, uid string) (*models.Student, error) {

	id, err := uuid.Parse(uid)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %w", err)
	}

	return s.repo.GetByID(ctx, id)
}

func (s *studentService) GetStudentByCardNumber(ctx context.Context, number string) (*models.Student, error) {
	return s.repo.GetByCardID(ctx, number)
}

func (s *studentService) GetAllStudents(ctx context.Context, params dto.PaginationParams) (*dto.StudentsResponse, error) {
	if params.Limit <= 0 {
		params.Limit = 10
	}
	if params.Limit > 100 {
		params.Limit = 100
	}
	if params.Offset < 0 {
		params.Offset = 0
	}

	students, total, err := s.repo.GetAll(ctx, params.Offset, params.Limit)
	if err != nil {
		return nil, err
	}

	pagination := dto.PaginationInfo{
		Offset:      params.Offset,
		Limit:       params.Limit,
		Total:       int(total),
		HasNext:     int64(params.Offset+params.Limit) < total,
		HasPrevious: params.Offset > 0,
	}

	response := &dto.StudentsResponse{
		Results:    students,
		Pagination: pagination,
	}

	return response, nil
}

func (s *studentService) DeleteStudent(ctx context.Context, uid string) error {
	id, err := uuid.Parse(uid)

	if err != nil {
		return fmt.Errorf("invalid uuid format: %w", err)
	}

	return s.repo.Delete(ctx, id)
}
