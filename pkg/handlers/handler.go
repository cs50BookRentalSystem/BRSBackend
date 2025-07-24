package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/getkin/kin-openapi/openapi3"

	"BRSBackend/pkg/api"
	"BRSBackend/pkg/services"
)

type Handler struct {
	bookService    services.BookService
	authService    services.AuthService
	studentService services.StudentService
	rentService    services.RentService
	reportService  services.ReportService
}

func NewHandler(svc *services.Service) *Handler {
	return &Handler{
		bookService:    svc.Book,
		authService:    svc.Auth,
		studentService: svc.Student,
		rentService:    svc.Rent,
		reportService:  svc.Report,
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (h *Handler) writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message, Code: statusCode})
}

func (h *Handler) writeResponse(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetOpenAPISpec(w http.ResponseWriter, r *http.Request) {
	swagger, err := api.GetSwagger()
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to get swagger: %v", err))
		return
	}
	headers := r.Header

	serverUrl := &url.URL{}
	if forwardedHost, ok := headers["X-Forwarded-Host"]; ok {
		proto := "http"
		if forwardedProto, ok := headers["X-Forwarded-Proto"]; ok {
			proto = forwardedProto[0]
		}
		serverUrl.Scheme = proto
		serverUrl.Host = forwardedHost[0]
	}

	if forwardedPrefix, ok := headers["X-Forwarded-Prefix"]; ok {
		serverUrl = serverUrl.JoinPath(forwardedPrefix[0])
	}

	if *serverUrl != (url.URL{}) {
		swagger.AddServer(&openapi3.Server{
			URL: serverUrl.String(),
		})
	}

	h.writeResponse(w, http.StatusOK, swagger)
}
