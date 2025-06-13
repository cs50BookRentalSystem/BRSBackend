package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"BRSBackend/pkg/dto"
	"BRSBackend/pkg/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func (h *AuthHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message, Code: statusCode})
}

func (h *AuthHandler) writeResponse(w http.ResponseWriter, statusCode int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if loginReq.User == "" || loginReq.Pass == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "Username and password are required")
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

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
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
