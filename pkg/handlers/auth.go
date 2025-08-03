package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/validation"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if validationErrors := validation.ValidateStruct(loginReq); validationErrors != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, validation.FormatErrors(validationErrors))
		return
	}

	response, sessionId, err := h.authService.Login(r.Context(), loginReq)
	if err != nil {
		h.writeErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)
	h.writeResponse(w, http.StatusOK, response)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		h.authService.Logout(r.Context(), cookie.Value)
	}

	clearCookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(-1 * time.Hour),
	}

	http.SetCookie(w, clearCookie)
	h.writeResponse(w, http.StatusOK, map[string]string{"message": "Logout successful"})
}

func (h *Handler) Librarian(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		h.writeErrorResponse(w, http.StatusUnauthorized, "invalid session or expired session")
		return
	}

	response, err := h.authService.GetLibrarian(r.Context(), cookie.Value)
	if err != nil {
		h.writeErrorResponse(w, http.StatusUnauthorized, "invalid session or expired session")
		return
	}

	h.writeResponse(w, http.StatusOK, response)
}
