package handlers

import (
	"context"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/models"
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

func (h *BookHandler) AddBook(ctx context.Context, request api.AddBookRequestObject) (api.AddBookResponseObject, error) {

	book := models.Book{
		Title:       request.Body.Title,
		Description: request.Body.Description,
		Count:       request.Body.Count,
	}

	if err, _ := h.bookService.CreateBook(ctx, &book); err != nil {
		return api.AddBook500JSONResponse{}, nil
	}

	return api.AddBook201JSONResponse{}, nil
}

//func main() {
//	port := flag.String("port", "8080", "Port for test HTTP server")
//	flag.Parse()
//
//	swagger, err := api.GetSwagger()
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
//		os.Exit(1)
//	}
//
//	// Clear out the servers array in the swagger spec, that skips validating
//	// that server names match. We don't know how this thing will be run.
//	swagger.Servers = nil
//	petStore := api.NewPetStore()
//
//	r := chi.NewRouter()
//	r.Use(middleware.OapiRequestValidator(swagger))
//	api.HandlerFromMux(petStore, r)
//
//	s := &http.Server{
//		Handler: r,
//		Addr:    net.JoinHostPort("0.0.0.0", *port),
//	}
//	log.Fatal(s.ListenAndServe())
//}
