package user

import "net/http"

// Sends a password reset token/link to user email
func (h *Handler) forgotPassword(w http.ResponseWriter, r *http.Request) {}

// Resets password using a verified reset token
func (h *Handler) resetPassword(w http.ResponseWriter, r *http.Request) {}

// Updates account password for an active logged-in user
func (h *Handler) updatePassword(w http.ResponseWriter, r *http.Request) {}
