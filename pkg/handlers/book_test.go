package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

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

func TestDeleteBook(t *testing.T) {
	t.Run("successful delete book", func(t *testing.T) {
		mockBookService := &services.MockBookService{
			DeleteBookFunc: func(ctx context.Context, id string) error {
				return nil
			},
		}

		h := NewHandler(&services.Service{Book: mockBookService})

		id := uuid.New()
		req := httptest.NewRequest(http.MethodDelete, "/books/"+id.String(), nil)
		w := httptest.NewRecorder()

		h.DeleteBookById(w, req, id)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("missing book id", func(t *testing.T) {
		mockBookService := &services.MockBookService{}
		h := NewHandler(&services.Service{Book: mockBookService})

		req := httptest.NewRequest(http.MethodDelete, "/books/", nil)
		w := httptest.NewRecorder()

		h.DeleteBookById(w, req, uuid.Nil)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("service error on delete", func(t *testing.T) {
		mockBookService := &services.MockBookService{
			DeleteBookFunc: func(ctx context.Context, id string) error {
				return errors.New("service error")
			},
		}

		h := NewHandler(&services.Service{Book: mockBookService})

		id := uuid.New()
		req := httptest.NewRequest(http.MethodDelete, "/books/"+id.String(), nil)
		w := httptest.NewRecorder()

		h.DeleteBookById(w, req, id)

		if w.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status code %d, got %d", http.StatusUnprocessableEntity, w.Code)
		}
	})
}
