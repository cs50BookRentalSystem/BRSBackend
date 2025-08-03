package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, string, error)
	ValidateSession(ctx context.Context, sessionId string) (*models.Librarian, error)
	Logout(ctx context.Context, sessionId string) error
	CreateLibrarian(ctx context.Context, username, password string) error
	CleanupExpiredSessions() error
	GetLibrarian(ctx context.Context, sessionId string) (*dto.LoginResponse, error)
}

type authService struct {
	librarianRepo repository.LibrarianRepository
	sessionRepo   repository.SessionRepository
}

func NewAuthService(librarianRepo repository.LibrarianRepository, sessionRepo repository.SessionRepository) AuthService {
	return &authService{
		librarianRepo: librarianRepo,
		sessionRepo:   sessionRepo,
	}
}

func (a *authService) GetLibrarian(ctx context.Context, sessionId string) (*dto.LoginResponse, error) {
	if sessionId == "" {
		return nil, errors.New("no session provided")
	}

	session, err := a.sessionRepo.GetByID(ctx, sessionId)
	if err != nil {
		return nil, fmt.Errorf("invalid session provided: %w", err)
	}

	response := &dto.LoginResponse{
		Message:     "valid session",
		LibrarianId: session.LibrarianId,
	}

	return response, nil
}

func (a *authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, string, error) {
	librarian, err := a.librarianRepo.GetByUsername(ctx, req.User)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword(librarian.Pass, []byte(req.Pass)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	sessionId, err := a.generateSessionID()
	if err != nil {
		return nil, "", errors.New("failed to create session")
	}

	session := &models.Session{
		Id:          sessionId,
		LibrarianId: librarian.Id,
		ExpiresAt:   time.Now().Add(24 * time.Hour),
	}

	if err := a.sessionRepo.Create(ctx, session); err != nil {
		return nil, "", fmt.Errorf("failed to create session: %w", err)
	}

	response := &dto.LoginResponse{
		Message:     "Login successful",
		LibrarianId: librarian.Id,
	}

	return response, sessionId, nil
}

func (a *authService) ValidateSession(ctx context.Context, sessionId string) (*models.Librarian, error) {
	if sessionId == "" {
		return nil, errors.New("no session provided")
	}

	session, err := a.sessionRepo.GetByID(ctx, sessionId)
	if err != nil {
		return nil, fmt.Errorf("invalid session provided: %w", err)
	}

	return &session.Librarian, nil
}

func (a *authService) Logout(ctx context.Context, sessionId string) error {
	if sessionId == "" {
		return nil
	}
	return a.sessionRepo.DeleteByID(ctx, sessionId)
}

func (a *authService) CreateLibrarian(ctx context.Context, username, password string) error {
	if _, err := a.librarianRepo.GetByUsername(ctx, username); err == nil {
		return errors.New("librarian already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	librarian := &models.Librarian{
		User: username,
		Pass: hashedPassword,
	}

	return a.librarianRepo.Create(ctx, librarian)
}

func (a *authService) generateSessionID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (a *authService) CleanupExpiredSessions() error {
	return a.sessionRepo.DeleteExpired()
}
