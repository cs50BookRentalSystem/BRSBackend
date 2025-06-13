package repository

import (
	"context"
	"time"

	"github.com/google/uuid"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
)

type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Book, error)
	GetAll(ctx context.Context, params dto.PaginationParams) ([]*models.Book, int64, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id uuid.UUID) error
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
	Update(ctx context.Context, cart *models.Cart) error
	UpdateStatus(ctx context.Context, cartID uuid.UUID, status string) error
}

type RentRepository interface {
	Create(ctx context.Context, rent *models.Rent) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Rent, error)
	GetByCartId(ctx context.Context, cartID uuid.UUID) ([]*models.Rent, error)
	GetAll(ctx context.Context, offset, limit int) ([]*models.Rent, int64, error)
	GetRentedBooks(ctx context.Context, studentCardID string) ([]*RentSummary, error)
	GetOverdueRentals(ctx context.Context, studentCardID string, offset, limit int) ([]*OverdueUser, int64, error)
	GetRentalReport(ctx context.Context) (*RentReport, error)
	SearchRents(ctx context.Context, bookName, studentName string, date *time.Time, offset, limit int) ([]*RentSummary, int64, error)
}

type SessionRepository interface {
	Create(ctx context.Context, session *models.Session) error
	GetByID(ctx context.Context, sessionId string) (*models.Session, error)
	DeleteByID(ctx context.Context, sessionId string) error
	DeleteExpired() error
	DeleteByLibrarianID(ctx context.Context, librarianId uuid.UUID) error
}

type RentSummary struct {
	RentID      uuid.UUID `json:"rent_id"`
	CartID      uuid.UUID `json:"cart_id"`
	BookTitle   string    `json:"book_title"`
	StudentName string    `json:"student_name"`
	RentedDate  time.Time `json:"rented_date"`
}

type OverdueUser struct {
	StudentName string    `json:"student_name"`
	Phone       string    `json:"phone"`
	TotalBooks  int       `json:"total_books"`
	DateRented  time.Time `json:"date_rented"`
	DaysOverdue int       `json:"days_overdue"`
}

type BookRentStats struct {
	BookTitle   string `json:"book_title"`
	RentedCount int    `json:"rented_count"`
}

type RentReport struct {
	TotalRents    int              `json:"total_rents"`
	TotalStudents int              `json:"total_students"`
	TopBooks      []*BookRentStats `json:"top_books"`
}

type Repository struct {
	Book      BookRepository
	Student   StudentRepository
	Librarian LibrarianRepository
	Cart      CartRepository
	Rent      RentRepository
	Session   SessionRepository
}
