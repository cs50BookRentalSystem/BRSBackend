package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/services"
)

func TestLogin(t *testing.T) {
	t.Run("successful login", func(t *testing.T) {
		mockAuthService := &services.MockAuthService{
			LoginFunc: func(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, string, error) {
				return &dto.LoginResponse{Message: "Login successful"}, "session-id", nil
			},
		}

		h := NewHandler(&services.Service{Auth: mockAuthService})

		body := dto.LoginRequest{User: "test", Pass: "password"}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		h.Login(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockAuthService := &services.MockAuthService{}
		h := NewHandler(&services.Service{Auth: mockAuthService})

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte("invalid json")))
		w := httptest.NewRecorder()

		h.Login(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("missing username or password", func(t *testing.T) {
		mockAuthService := &services.MockAuthService{}
		h := NewHandler(&services.Service{Auth: mockAuthService})

		body := dto.LoginRequest{User: "", Pass: ""}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		h.Login(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestLibrarian(t *testing.T) {
	t.Run("successful get librarian", func(t *testing.T) {
		mockAuthService := &services.MockAuthService{
			GetLibrarianFunc: func(ctx context.Context, sessionID string) (*dto.LoginResponse, error) {
				return &dto.LoginResponse{LibrarianId: uuid.New(), Message: "test"}, nil
			},
		}

		h := NewHandler(&services.Service{Auth: mockAuthService})

		req := httptest.NewRequest(http.MethodGet, "/librarian", nil)
		req.AddCookie(&http.Cookie{Name: "session_id", Value: "session-id"})
		w := httptest.NewRecorder()

		h.Librarian(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("invalid session", func(t *testing.T) {
		mockAuthService := &services.MockAuthService{
			GetLibrarianFunc: func(ctx context.Context, sessionID string) (*dto.LoginResponse, error) {
				return nil, fmt.Errorf("invalid session")
			},
		}

		h := NewHandler(&services.Service{Auth: mockAuthService})

		req := httptest.NewRequest(http.MethodGet, "/librarian", nil)
		req.AddCookie(&http.Cookie{Name: "session_id", Value: "session-id"})
		w := httptest.NewRecorder()

		h.Librarian(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("expected status code %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})

	t.Run("missing session cookie", func(t *testing.T) {
		mockAuthService := &services.MockAuthService{}
		h := NewHandler(&services.Service{Auth: mockAuthService})

		req := httptest.NewRequest(http.MethodGet, "/librarian", nil)
		w := httptest.NewRecorder()

		h.Librarian(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("expected status code %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})
}
