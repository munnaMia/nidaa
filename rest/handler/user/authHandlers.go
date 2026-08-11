package user

import (
	"errors"
	"net/http"

	"github.com/munnaMia/nidaa/internal/domain"
)

// Registers a new user account
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var req registerRequest

	// decode the req
	if err := h.jshlp.Decoder(r.Body, &req); err != nil {
		h.responder.SendError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// validate req
	if ok := checkRegisterReq(w, h.validate, h.responder, req); !ok {
		return
	}

	token, user, err := h.uc.RegisterUser(r.Context(), req.UserName, req.Name, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrEmailAlreadyExist) || errors.Is(err, domain.ErrUsernameAlreadyExist) {
			h.responder.SendError(w, http.StatusConflict, err.Error())
			return
		}
		h.responder.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if user == nil {
		h.responder.SendError(w, http.StatusInternalServerError, "User creation failed")
		return
	}

	// Construct a response
	regRes := &authResponse{
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
func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

	err := h.jshlp.Decoder(r.Body, &req)
	if err != nil {
		h.responder.SendError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// validate login req
	if ok := checkLoginReq(w, h.validate, h.responder, req); !ok {
		return
	}

	jwt, user, err := h.uc.LoginUser(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			h.responder.SendError(w, http.StatusUnauthorized, err.Error())
			return
		}
		h.responder.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	logRes := &authResponse{
		Token: jwt,
		User: &userResponse{
			UserName: user.UserName,
			Name:     user.Name,
			Email:    user.Email,
		},
	}

	h.responder.SendResponse(w, http.StatusOK, logRes, nil)
}

// Clears cookies/tokens and invalidates current session
func (h *Handler) logoutUser(w http.ResponseWriter, r *http.Request) {}

// Issues a new access token using a valid refresh token
func (h *Handler) refreshToken(w http.ResponseWriter, r *http.Request) {}
