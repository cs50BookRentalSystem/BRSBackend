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

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) repository.StudentRepository {
	return &studentRepository{db: db}
}

func (s studentRepository) Create(ctx context.Context, student *models.Student) error {
	if err := s.db.WithContext(ctx).Create(student).Error; err != nil {
		return fmt.Errorf("failed to create student: %w", err)
	}

	return nil
}

func (s studentRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Student, error) {
	var student models.Student
	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("student not found")
		}
		return nil, fmt.Errorf("failed to get student: %w", err)
	}

	return &student, nil
}

func (s studentRepository) GetByCardID(ctx context.Context, cardID string) (*models.Student, error) {
	var student models.Student
	if err := s.db.WithContext(ctx).Where("card_id = ?", cardID).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("student not found")
		}
		return nil, fmt.Errorf("failed to get student: %w", err)
	}

	return &student, nil
}

func (s studentRepository) GetAll(ctx context.Context, offset, limit int) ([]*models.Student, int64, error) {
	var students []*models.Student
	var total int64

	if err := s.db.WithContext(ctx).
		Model(&models.Student{}).
		Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count students: %w", err)
	}

	if err := s.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&students).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get students: %w", err)
	}

	return students, total, nil
}

func (s studentRepository) Update(ctx context.Context, student *models.Student) error {
	panic("implement me")
}

func (s studentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Student{}).Error
}
