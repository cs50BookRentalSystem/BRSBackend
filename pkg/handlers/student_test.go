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

func TestAddStudent(t *testing.T) {
	t.Run("successful add student", func(t *testing.T) {
		mockStudentService := &services.MockStudentService{
			CreateStudentFunc: func(ctx context.Context, student *models.Student) error {
				return nil
			},
		}

		h := NewHandler(&services.Service{Student: mockStudentService})

		body := models.Student{FirstName: "Test", LastName: "User", Major: "CS", Phone: "1234567890"}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		h.AddStudent(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockStudentService := &services.MockStudentService{}
		h := NewHandler(&services.Service{Student: mockStudentService})

		req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader([]byte("invalid json")))
		w := httptest.NewRecorder()

		h.AddStudent(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("missing first name", func(t *testing.T) {
		mockStudentService := &services.MockStudentService{}
		h := NewHandler(&services.Service{Student: mockStudentService})

		body := models.Student{LastName: "User", Major: "CS", Phone: "1234567890"}
		bodyBytes, _ := json.Marshal(body)

		req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewReader(bodyBytes))
		w := httptest.NewRecorder()

		h.AddStudent(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}
