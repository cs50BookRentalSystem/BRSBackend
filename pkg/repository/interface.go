package repository

import (
	"context"

	"github.com/google/uuid"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
)

type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Book, error)
	GetAll(ctx context.Context, params dto.PaginationParams) ([]*models.Book, int64, error)
	GetBooksByIDs(ctx context.Context, bookIDs []uuid.UUID) ([]*models.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
	UpdateCount(ctx context.Context, bookID uuid.UUID, delta int) error
	DecrementCount(ctx context.Context, bookID uuid.UUID) error
	IncrementCount(ctx context.Context, bookID uuid.UUID) error
	DecrementMultipleBooks(ctx context.Context, bookIDs []uuid.UUID) error
	IncrementMultipleBooks(ctx context.Context, bookIDs []uuid.UUID) error
}

type StudentRepository interface {
	Create(ctx context.Context, student *models.Student) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Student, error)
	GetByCardID(ctx context.Context, cardID string) (*models.Student, error)
	GetAll(ctx context.Context, offset, limit int) ([]*models.Student, int64, error)
	Update(ctx context.Context, student *models.Student) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type LibrarianRepository interface {
	Create(ctx context.Context, librarian *models.Librarian) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Librarian, error)
	GetByUsername(ctx context.Context, username string) (*models.Librarian, error)
}

type CartRepository interface {
	Create(ctx context.Context, cart *models.Cart) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Cart, error)
	GetByStatus(ctx context.Context, status string) ([]*models.Cart, error)
	GetCartsByStudentID(ctx context.Context, studentID uuid.UUID) ([]*models.Cart, error)
	UpdateStatus(ctx context.Context, cartID uuid.UUID, status string) error
}

type RentRepository interface {
	Create(ctx context.Context, rent *models.Rent) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Rent, error)
	GetRentsByFilters(ctx context.Context, filters dto.RentFilters) ([]*dto.RentSummary, int64, error)
	GetRentedBooksByStudent(ctx context.Context, studentCardID string) ([]*dto.RentSummary, error)
	GetRentsByCartID(ctx context.Context, cartID uuid.UUID) ([]*models.Rent, error)
}

type SessionRepository interface {
	Create(ctx context.Context, session *models.Session) error
	GetByID(ctx context.Context, sessionId string) (*models.Session, error)
	DeleteByID(ctx context.Context, sessionId string) error
	DeleteExpired() error
	DeleteByLibrarianID(ctx context.Context, librarianId uuid.UUID) error
}

type ReportRepository interface {
	GetOverdueRentals(ctx context.Context, studentCardID *string, limit, offset, overduePeriod int) ([]dto.OverdueUser, int64, error)
	GetRentalReport(ctx context.Context, limit, offset, overduePeriod int) (*dto.RentReport, error)
}

type Repository struct {
	Book      BookRepository
	Student   StudentRepository
	Librarian LibrarianRepository
	Cart      CartRepository
	Rent      RentRepository
	Session   SessionRepository
	Report    ReportRepository
}
