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

func TestDeleteStudent(t *testing.T) {
	t.Run("successful delete student", func(t *testing.T) {
		mockStudentService := &services.MockStudentService{
			DeleteStudentFunc: func(ctx context.Context, id string) error {
				return nil
			},
		}

		h := NewHandler(&services.Service{Student: mockStudentService})

		id := uuid.New()
		req := httptest.NewRequest(http.MethodDelete, "/students/"+id.String(), nil)
		w := httptest.NewRecorder()

		h.DeleteStudentById(w, req, id)

		if w.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("missing student id", func(t *testing.T) {
		mockStudentService := &services.MockStudentService{}
		h := NewHandler(&services.Service{Student: mockStudentService})

		req := httptest.NewRequest(http.MethodDelete, "/students/", nil)
		w := httptest.NewRecorder()

		h.DeleteStudentById(w, req, uuid.Nil)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("service error on delete", func(t *testing.T) {
		mockStudentService := &services.MockStudentService{
			DeleteStudentFunc: func(ctx context.Context, id string) error {
				return errors.New("service error")
			},
		}

		h := NewHandler(&services.Service{Student: mockStudentService})

		id := uuid.New()
		req := httptest.NewRequest(http.MethodDelete, "/students/"+id.String(), nil)
		w := httptest.NewRecorder()

		h.DeleteStudentById(w, req, id)

		if w.Code != http.StatusUnprocessableEntity {
			t.Errorf("expected status code %d, got %d", http.StatusUnprocessableEntity, w.Code)
		}
	})
}
