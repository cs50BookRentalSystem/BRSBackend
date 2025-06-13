package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"BRSBackend/pkg/services"
)

type contextKey string

const LibrarianContextKey contextKey = "librarian"

func AuthMiddleware(authService services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if isPublicRoute(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie("session_id")
			if err != nil {
				writeUnauthorizedResponse(w, "No session cookie found")
				return
			}

			librarian, err := authService.ValidateSession(r.Context(), cookie.Value)
			if err != nil {
				writeUnauthorizedResponse(w, "Invalid session")
				return
			}

			ctx := context.WithValue(r.Context(), LibrarianContextKey, librarian)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func isPublicRoute(path string) bool {
	publicRoutes := []string{
		"/login",
		"/logout",
	}

	for _, route := range publicRoutes {
		if path == route {
			return true
		}
	}
	return false
}

func writeUnauthorizedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
