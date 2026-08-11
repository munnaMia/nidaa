package user

import "net/http"

// Retrieves profile info for the logged-in user
func (h *Handler) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	r.Context().Value()
}

// Updates partial profile data (e.g., name, avatar, bio)
func (h *Handler) updateCurrentUser(w http.ResponseWriter, r *http.Request) {}

// Deletes or deactivates the current user account
func (h *Handler) deleteCurrentUser(w http.ResponseWriter, r *http.Request) {}
