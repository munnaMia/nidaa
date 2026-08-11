package user

import (
	"errors"
	"net/http"

	"github.com/munnaMia/nidaa/internal/domain"
	reqctx "github.com/munnaMia/nidaa/internal/reqCtx"
)

// Retrieves profile info for the logged-in user
func (h *Handler) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	id, ok := reqctx.GetUserID(r.Context())
	if !ok {
		h.responder.SendError(w, http.StatusUnauthorized, "unauthorized access.")
		return
	}

	user, err := h.uc.GetUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			h.responder.SendError(w, http.StatusNotFound, err.Error())
			return
		}
		h.responder.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	userResponse := &userResponse{
		Name:     user.Name,
		UserName: user.UserName,
		Email:    user.Email,
	}

	h.responder.SendResponse(w, http.StatusOK, userResponse, nil)

}

// Updates partial profile data (e.g., name, avatar, bio)
func (h *Handler) updateCurrentUser(w http.ResponseWriter, r *http.Request) {}

// Deletes or deactivates the current user account
func (h *Handler) deleteCurrentUser(w http.ResponseWriter, r *http.Request) {}
