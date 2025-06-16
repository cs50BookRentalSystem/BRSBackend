package middleware

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3filter"

	"BRSBackend/pkg/services"
)

type contextKey string

const LibrarianContextKey contextKey = "librarian"

func NewOApiAuthenticationFunc(authService services.AuthService) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		req := input.RequestValidationInput.Request

		cookie, err := req.Cookie("session_id")
		if err != nil {
			return fmt.Errorf("auth failed: %w", err)
		}

		librarian, err := authService.ValidateSession(ctx, cookie.Value)
		if err != nil {
			return fmt.Errorf("session validation failed: %w", err)
		}

		newCtx := context.WithValue(req.Context(), LibrarianContextKey, librarian)
		input.RequestValidationInput.Request = req.WithContext(newCtx)
		return nil
	}
}
