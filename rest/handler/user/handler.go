package user

import "net/http"

type Handler struct {
}

// create a user handler
func NewHandler() *Handler {
	return &Handler{}
}

// Registers a new user account
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {}

// Authenticates user; returns JWT access/refresh tokens
func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {}

// Clears cookies/tokens and invalidates current session
func (h *Handler) logoutUser(w http.ResponseWriter, r *http.Request) {}

// Issues a new access token using a valid refresh token
func (h *Handler) refreshToken(w http.ResponseWriter, r *http.Request) {}

// Sends a password reset token/link to user email
func (h *Handler) forgotPassword(w http.ResponseWriter, r *http.Request) {}

// Resets password using a verified reset token
func (h *Handler) resetPassword(w http.ResponseWriter, r *http.Request) {}

// Updates account password for an active logged-in user
func (h *Handler) updatePassword(w http.ResponseWriter, r *http.Request) {}

// Retrieves profile info for the logged-in user
func (h *Handler) getCurrentUser(w http.ResponseWriter, r *http.Request) {}

// Updates partial profile data (e.g., name, avatar, bio)
func (h *Handler) updateCurrentUser(w http.ResponseWriter, r *http.Request) {}

// Deletes or deactivates the current user account
func (h *Handler) deleteCurrentUser(w http.ResponseWriter, r *http.Request) {}
