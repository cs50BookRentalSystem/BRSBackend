package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"BRSBackend/pkg/models"
	"BRSBackend/pkg/services"
)

func TestAddBook(t *testing.T) {
	t.Run("successful add book", func(t *testing.T) {
		mockBookService := &services.MockBookService{
			CreateBookFunc: func(ctx context.Context, book *models.Book) error {
				return nil
			},
		}

		h := NewHandler(&services.Service{Book: mockBookService})

		body := models.Book{Title: "Test Book", Count: 10}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		h.AddBook(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockBookService := &services.MockBookService{}
		h := NewHandler(&services.Service{Book: mockBookService})

		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader([]byte("invalid json")))
		w := httptest.NewRecorder()

		h.AddBook(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("missing title", func(t *testing.T) {
		mockBookService := &services.MockBookService{}
		h := NewHandler(&services.Service{Book: mockBookService})

		body := models.Book{Count: 10}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		h.AddBook(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}
