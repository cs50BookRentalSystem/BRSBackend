package sqlite

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) repository.SessionRepository {
	return &sessionRepository{db: db}
}

func (s *sessionRepository) Create(ctx context.Context, session *models.Session) error {
	if err := s.db.WithContext(ctx).Create(session).Error; err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	return nil
}

func (s *sessionRepository) GetByID(ctx context.Context, sessionId string) (*models.Session, error) {
	var session models.Session
	if err := s.db.WithContext(ctx).Where("id = ? AND expires_at > ?", sessionId, time.Now()).
		Preload("Librarian").First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("session not found")
		}
		return nil, fmt.Errorf("failed to get session id: %w", err)
	}

	return &session, nil
}

func (s *sessionRepository) DeleteByID(ctx context.Context, sessionId string) error {
	return s.db.WithContext(ctx).Where("id = ?", sessionId).Delete(&models.Session{}).Error
}

func (s *sessionRepository) DeleteExpired() error {
	return s.db.Where("expires_at <= ?", time.Now()).Delete(&models.Session{}).Error
}

func (s *sessionRepository) DeleteByLibrarianID(ctx context.Context, librarianId uuid.UUID) error {
	return s.db.WithContext(ctx).Where("librarian_id = ?", librarianId).Delete(&models.Session{}).Error
}
