package user

import (
	"net/http"

	"github.com/munnaMia/nidaa/rest/middleware"
)

func (h *Handler) RegisterRoute(mux *http.ServeMux, mngr *middleware.Manager) {

	// Authentication Routes
	mux.Handle("POST /api/auth/register", mngr.With(
		http.HandlerFunc(h.registerUser),
	))
	mux.Handle("POST /api/auth/login", mngr.With(
		http.HandlerFunc(h.loginUser),
	))
	mux.Handle("POST /api/auth/logout", mngr.With(
		http.HandlerFunc(h.logoutUser),
	))
	mux.Handle("POST /api/auth/refresh-token", mngr.With(
		http.HandlerFunc(h.refreshToken),
	))

	// Password Security & Recovery
	mux.Handle("POST /api/auth/passwords/forgot", mngr.With(
		http.HandlerFunc(h.forgotPassword),
	))
	mux.Handle("POST /api/auth/passwords/reset", mngr.With(
		http.HandlerFunc(h.resetPassword),
	))
	mux.Handle("PUT /api/account/me/password", mngr.With(
		http.HandlerFunc(h.updatePassword),
	))

	// Account & Profile Routes
	mux.Handle("GET /api/account/me", mngr.With(
		http.HandlerFunc(h.getCurrentUser),
	))
	mux.Handle("PATCH /api/account/me", mngr.With(
		http.HandlerFunc(h.updateCurrentUser),
	))
	mux.Handle("DELETE /api/account/me", mngr.With(
		http.HandlerFunc(h.deleteCurrentUser),
	))
}
