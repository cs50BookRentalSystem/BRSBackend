package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/models"
	"BRSBackend/pkg/repository"
)

type RentService interface {
	CreateRentTransaction(ctx context.Context, req dto.CreateRentRequest) (*dto.CreateRentResponse, error)
	GetRents(ctx context.Context, filters dto.RentFilters) (*dto.GetRentedBooksResponse, error)
	GetRentedBooksByStudent(ctx context.Context, studentCardID *string) (*dto.GetRentedBooksResponse, error)
	ReturnBooks(ctx context.Context, cartID uuid.UUID) (*dto.ReturnBooksResponse, error)
}

type rentService struct {
	rentRepo    repository.RentRepository
	cartRepo    repository.CartRepository
	bookRepo    repository.BookRepository
	studentRepo repository.StudentRepository
}

func NewRentService(
	rentRepo repository.RentRepository,
	cartRepo repository.CartRepository,
	bookRepo repository.BookRepository,
	studentRepo repository.StudentRepository,
) RentService {
	return &rentService{
		rentRepo:    rentRepo,
		cartRepo:    cartRepo,
		bookRepo:    bookRepo,
		studentRepo: studentRepo,
	}
}

func (r *rentService) CreateRentTransaction(ctx context.Context, req dto.CreateRentRequest) (*dto.CreateRentResponse, error) {
	student, err := r.studentRepo.GetByID(ctx, req.StudentID)
	if err != nil {
		return nil, fmt.Errorf("student not found: %w", err)
	}
	if student == nil {
		return nil, fmt.Errorf("student with ID %s not found", req.StudentID)
	}

	books, err := r.bookRepo.GetBooksByIDs(ctx, req.BookIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch books: %w", err)
	}
	if len(books) != len(req.BookIDs) {
		return nil, fmt.Errorf("one or more books not found")
	}

	cart := &models.Cart{
		StudentId: req.StudentID,
		Status:    "RENTED",
	}

	if err := r.cartRepo.Create(ctx, cart); err != nil {
		return nil, fmt.Errorf("failed to create cart: %w", err)
	}

	for _, bookID := range req.BookIDs {
		rent := &models.Rent{
			CartId: cart.Id,
			BookId: bookID,
		}

		if err := r.rentRepo.Create(ctx, rent); err != nil {
			return nil, fmt.Errorf("failed to create rent record for book %s: %w", bookID, err)
		}
	}

	return &dto.CreateRentResponse{
		CartID:  cart.Id,
		Message: "Books rented successfully",
	}, nil
}

func (r *rentService) GetRents(ctx context.Context, filters dto.RentFilters) (*dto.GetRentedBooksResponse, error) {

	if filters.Limit <= 0 {
		filters.Limit = 10
	}

	if filters.Offset < 0 {
		filters.Offset = 0
	}

	repoFilters := dto.RentFilters{
		BookName:    filters.BookName,
		StudentName: filters.StudentName,
		Date:        filters.Date,
		Limit:       filters.Limit,
		Offset:      filters.Offset,
	}

	rents, err := r.rentRepo.GetRentsByFilters(ctx, repoFilters)
	if err != nil {
		return nil, fmt.Errorf("failed to get rents: %w", err)
	}

	if filters.Limit > 100 {
		filters.Limit = 100
	}

	pagination := dto.PaginationInfo{
		Total:       len(rents),
		Limit:       filters.Limit,
		Offset:      filters.Offset,
		HasNext:     filters.Offset+filters.Limit < len(rents),
		HasPrevious: filters.Offset > 0,
	}

	return &dto.GetRentedBooksResponse{
		Results:    rents,
		Pagination: pagination,
	}, nil
}

func (r *rentService) GetRentedBooksByStudent(ctx context.Context, studentCardID *string) (*dto.GetRentedBooksResponse, error) {
	cardID := ""
	if studentCardID != nil {
		cardID = *studentCardID
	}

	rents, err := r.rentRepo.GetRentedBooksByStudent(ctx, cardID)
	if err != nil {
		return nil, fmt.Errorf("failed to get rented books: %w", err)
	}

	pagination := dto.PaginationInfo{
		Total:  len(rents),
		Limit:  0,
		Offset: 0,
	}

	return &dto.GetRentedBooksResponse{
		Results:    rents,
		Pagination: pagination,
	}, nil
}

func (r *rentService) ReturnBooks(ctx context.Context, cartID uuid.UUID) (*dto.ReturnBooksResponse, error) {
	cart, err := r.cartRepo.GetByID(ctx, cartID)
	if err != nil {
		return nil, fmt.Errorf("cart not found: %w", err)
	}

	if cart.Status != "RENTED" {
		return nil, fmt.Errorf("cart %s is not currently rented (status: %s)", cartID, cart.Status)
	}

	if err := r.cartRepo.UpdateStatus(ctx, cartID, "RETURNED"); err != nil {
		return nil, fmt.Errorf("failed to update cart status: %w", err)
	}

	return &dto.ReturnBooksResponse{
		Message: "Cart marked as returned",
		CartID:  cartID,
	}, nil
}
