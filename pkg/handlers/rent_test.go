package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/services"
)

func TestCreateRentTransaction(t *testing.T) {
	t.Run("successful rent creation", func(t *testing.T) {
		mockRentService := &services.MockRentService{
			CreateRentTransactionFunc: func(ctx context.Context, req dto.CreateRentRequest) (*dto.CreateRentResponse, error) {
				return &dto.CreateRentResponse{Message: "Rent created successfully"}, nil
			},
		}

		h := NewHandler(&services.Service{Rent: mockRentService})

		body := dto.CreateRentRequest{StudentID: uuid.New(), BookIDs: []uuid.UUID{uuid.New()}}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/rents", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		h.CreateRentTransaction(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockRentService := &services.MockRentService{}
		h := NewHandler(&services.Service{Rent: mockRentService})

		req := httptest.NewRequest(http.MethodPost, "/rents", bytes.NewReader([]byte("invalid json")))
		w := httptest.NewRecorder()

		h.CreateRentTransaction(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("missing student id", func(t *testing.T) {
		mockRentService := &services.MockRentService{}
		h := NewHandler(&services.Service{Rent: mockRentService})

		body := dto.CreateRentRequest{BookIDs: []uuid.UUID{uuid.New()}}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/rents", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		h.CreateRentTransaction(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}
