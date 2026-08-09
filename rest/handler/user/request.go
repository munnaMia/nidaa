package user

import (
	"net/http"

	"github.com/munnaMia/nidaa/util/responder"
	"github.com/munnaMia/nidaa/util/validate"
)

type registerRequest struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// default rule
var defaultPasswordRule = validate.PasswordRules{
	MaxLength:      100,
	MinLenght:      8,
	RequireUpper:   false,
	RequireLower:   false,
	RequireNumber:  false,
	RequireSpecial: false,
}

// check register req. if it is a validated req then it return true and if not a valid req then it return http response and return false.
func checkRegisterReq(w http.ResponseWriter, validate validate.Validator, responder responder.Responder, req registerRequest) bool {
	if err := validate.String(50, 8, req.UserName); err != nil {
		responder.SendError(w, http.StatusBadRequest, err.Error())
		return false
	}

	if err := validate.String(255, 8, req.Name); err != nil {
		responder.SendError(w, http.StatusBadRequest, err.Error())
		return false
	}

	if err := validate.Email(req.Email); err != nil {
		responder.SendError(w, http.StatusBadRequest, err.Error())
		return false
	}

	if errs := validate.Password(req.Password, defaultPasswordRule); errs != nil {
		responder.SendError(w, http.StatusBadRequest, "invalid password", errs...)
		return false
	}

	return true
}

// check login req. if it is a validated req then it return true and if not a valid req then it return http response and return false.
func checkLoginReq(w http.ResponseWriter, validate validate.Validator, responder responder.Responder, req loginRequest) bool {
	if err := validate.Email(req.Email); err != nil {
		responder.SendError(w, http.StatusBadRequest, err.Error())
		return false
	}

	if errs := validate.Password(req.Password, defaultPasswordRule); errs != nil {
		responder.SendError(w, http.StatusBadRequest, "invalid password", errs...)
		return false
	}

	return true
}
