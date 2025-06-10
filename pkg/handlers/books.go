package handlers

import (
	_ "context"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/models"
	_ "BRSBackend/pkg/repository"
	"BRSBackend/pkg/repository/sqlite"
	"BRSBackend/pkg/services"
)

type BookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

type Error struct {
	Code    int32
	Message string
}

func sendLibError(w http.ResponseWriter, code int, message string) {
	petErr := Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(petErr)
}

func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {

	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		return
	}

	if err, _ := h.bookService.CreateBook(r.Context(), &newBook); err != nil {
		return
	}

	return
}

func (h *BookHandler) ListOrSearchBooks(w http.ResponseWriter, r *http.Request, params api.ListOrSearchBooksParams) {
	panic("implement me")
}

// librarian login
// (POST /login)
func (h *BookHandler) Login(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// Logout librarian
// (POST /logout)
func (h *BookHandler) Logout(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// Get overdue rentals
// (GET /overdues)
func (h *BookHandler) ListOverdueRentals(w http.ResponseWriter, r *http.Request, params api.ListOverdueRentalsParams) {
	panic("implement me")
}

// Get list of all rents with optional filters
// (GET /rents)
func (h *BookHandler) ListRents(w http.ResponseWriter, r *http.Request, params api.ListRentsParams) {
	panic("implement me")
}

// Create rental transaction
// (POST /rents)
func (h *BookHandler) CreateRentTransaction(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// Get rental report
// (GET /reports)
func (h *BookHandler) GetRentalReports(w http.ResponseWriter, r *http.Request, params api.GetRentalReportsParams) {
	panic("implement me")
}

// List books currently rented by a student
// (GET /returns)
func (h *BookHandler) GetRentedBooksByStudent(w http.ResponseWriter, r *http.Request, params api.GetRentedBooksByStudentParams) {
	panic("implement me")
}

// Mark a cart as returned
// (PUT /returns)
func (h *BookHandler) ReturnBooks(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// List all students
// (GET /students)
func (h *BookHandler) ListAllStudents(w http.ResponseWriter, r *http.Request, params api.ListAllStudentsParams) {
	panic("implement me")
}

// Register a new student
// (POST /students)
func (h *BookHandler) AddStudent(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// Get a Student by ID
// (GET /students/{id})
func (h *BookHandler) GetStudentById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	panic("implement me")
}

func main() {
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	a := sqlite.NewBookRepository(&gorm.DB{})

	b := services.NewBookService(a)
	petStore := NewBookHandler(b)

	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidator(swagger))
	api.HandlerFromMux(petStore, r)

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", *port),
	}
	err = s.ListenAndServe()
	if err != nil {
		return
	}
}
