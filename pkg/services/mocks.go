package services

import (
	"context"

	"github.com/google/uuid"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
)

type MockAuthService struct {
	LoginFunc                  func(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, string, error)
	ValidateSessionFunc        func(ctx context.Context, sessionId string) (*models.Librarian, error)
	LogoutFunc                 func(ctx context.Context, sessionId string) error
	CreateLibrarianFunc        func(ctx context.Context, username, password string) error
	CleanupExpiredSessionsFunc func() error
	GetLibrarianFunc           func(ctx context.Context, sessionID string) (*dto.LoginResponse, error)
}

func (m *MockAuthService) GetLibrarian(ctx context.Context, sessionId string) (*dto.LoginResponse, error) {
	return m.GetLibrarianFunc(ctx, sessionId)
}

func (m *MockAuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, string, error) {
	return m.LoginFunc(ctx, req)
}

func (m *MockAuthService) ValidateSession(ctx context.Context, sessionId string) (*models.Librarian, error) {
	return m.ValidateSessionFunc(ctx, sessionId)
}

func (m *MockAuthService) Logout(ctx context.Context, sessionId string) error {
	return m.LogoutFunc(ctx, sessionId)
}

func (m *MockAuthService) CreateLibrarian(ctx context.Context, username, password string) error {
	return m.CreateLibrarianFunc(ctx, username, password)
}

func (m *MockAuthService) CleanupExpiredSessions() error {
	return m.CleanupExpiredSessionsFunc()
}

type MockBookService struct {
	CreateBookFunc  func(ctx context.Context, book *models.Book) error
	GetBookByIDFunc func(ctx context.Context, id string) (*models.Book, error)
	GetAllBooksFunc func(ctx context.Context, params dto.PaginationParams) (*dto.BooksResponse, error)
	DeleteBookFunc  func(ctx context.Context, id string) error
}

func (m *MockBookService) CreateBook(ctx context.Context, book *models.Book) error {
	return m.CreateBookFunc(ctx, book)
}

func (m *MockBookService) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	return m.GetBookByIDFunc(ctx, id)
}

func (m *MockBookService) GetAllBooks(ctx context.Context, params dto.PaginationParams) (*dto.BooksResponse, error) {
	return m.GetAllBooksFunc(ctx, params)
}

func (m *MockBookService) DeleteBook(ctx context.Context, id string) error {
	return m.DeleteBookFunc(ctx, id)
}

type MockStudentService struct {
	CreateStudentFunc          func(ctx context.Context, student *models.Student) error
	GetStudentByIDFunc         func(ctx context.Context, id string) (*models.Student, error)
	GetStudentByCardNumberFunc func(ctx context.Context, number string) (*models.Student, error)
	GetAllStudentsFunc         func(ctx context.Context, params dto.PaginationParams) (*dto.StudentsResponse, error)
	DeleteStudentFunc          func(ctx context.Context, id string) error
}

func (m *MockStudentService) DeleteStudent(ctx context.Context, id string) error {
	return m.DeleteStudentFunc(ctx, id)
}

func (m *MockStudentService) CreateStudent(ctx context.Context, student *models.Student) error {
	return m.CreateStudentFunc(ctx, student)
}

func (m *MockStudentService) GetStudentByID(ctx context.Context, id string) (*models.Student, error) {
	return m.GetStudentByIDFunc(ctx, id)
}

func (m *MockStudentService) GetStudentByCardNumber(ctx context.Context, number string) (*models.Student, error) {
	return m.GetStudentByCardNumberFunc(ctx, number)
}

func (m *MockStudentService) GetAllStudents(ctx context.Context, params dto.PaginationParams) (*dto.StudentsResponse, error) {
	return m.GetAllStudentsFunc(ctx, params)
}

type MockRentService struct {
	CreateRentTransactionFunc   func(ctx context.Context, req dto.CreateRentRequest) (*dto.CreateRentResponse, error)
	GetRentsFunc                func(ctx context.Context, filters dto.RentFilters) (*dto.GetRentedBooksResponse, error)
	GetRentedBooksByStudentFunc func(ctx context.Context, studentCardID *string) (*dto.GetRentedBooksResponse, error)
	ReturnBooksFunc             func(ctx context.Context, cartID uuid.UUID) (*dto.ReturnBooksResponse, error)
}

func (m *MockRentService) CreateRentTransaction(ctx context.Context, req dto.CreateRentRequest) (*dto.CreateRentResponse, error) {
	return m.CreateRentTransactionFunc(ctx, req)
}

func (m *MockRentService) GetRents(ctx context.Context, filters dto.RentFilters) (*dto.GetRentedBooksResponse, error) {
	return m.GetRentsFunc(ctx, filters)
}

func (m *MockRentService) GetRentedBooksByStudent(ctx context.Context, studentCardID *string) (*dto.GetRentedBooksResponse, error) {
	return m.GetRentedBooksByStudentFunc(ctx, studentCardID)
}

func (m *MockRentService) ReturnBooks(ctx context.Context, cartID uuid.UUID) (*dto.ReturnBooksResponse, error) {
	return m.ReturnBooksFunc(ctx, cartID)
}

type MockReportService struct {
	GetOverdueRentalsFunc func(ctx context.Context, studentCardID *string, limit, offset int) (*dto.OverdueResponse, error)
	GetRentalReportFunc   func(ctx context.Context, limit, offset int) (*dto.RentReport, error)
}

func (m *MockReportService) GetOverdueRentals(ctx context.Context, studentCardID *string, limit, offset int) (*dto.OverdueResponse, error) {
	return m.GetOverdueRentalsFunc(ctx, studentCardID, limit, offset)
}

func (m *MockReportService) GetRentalReport(ctx context.Context, limit, offset int) (*dto.RentReport, error) {
	return m.GetRentalReportFunc(ctx, limit, offset)
}
