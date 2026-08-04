package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/munnaMia/nidaa/internal/domain"
	"github.com/munnaMia/nidaa/internal/usecase"
	"github.com/munnaMia/nidaa/util/responder"
)

type Handler struct {
	uc        *usecase.UserUseCase
	responder responder.Responder
}

// create a user handler
func NewHandler(
	uc *usecase.UserUseCase,
	res responder.Responder,
) *Handler {
	return &Handler{
		uc:        uc,
		responder: res,
	}
}

// Registers a new user account
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	// 1. Limit request body size (e.g., 1MB) to prevent DoS attacks
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024)

	var req registerRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&req)
	if err != nil {
		h.responder.SendError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	token, user, err := h.uc.RegisterUser(r.Context(), req.UserName, req.Name, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrEmailAlreadyExist) {
			h.responder.SendError(w, http.StatusConflict, err.Error())
		} else if errors.Is(err, domain.ErrUsernameAlreadyExist) {
			h.responder.SendError(w, http.StatusConflict, err.Error())
		} else {
			h.responder.SendError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	if user == nil {
		h.responder.SendError(w, http.StatusInternalServerError, "User creation failed")
		return
	}

	// Construct a response
	regRes := &registerResponse{
		Token: token,
		User: &userResponse{
			Name:     user.Name,
			UserName: user.UserName,
			Email:    user.Email,
		},
	}

	h.responder.SendResponse(w, http.StatusCreated, regRes, nil)

}

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
