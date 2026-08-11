package user

import (
	"github.com/munnaMia/nidaa/internal/usecase"
	jsonhelper "github.com/munnaMia/nidaa/util/jsonHelper"
	"github.com/munnaMia/nidaa/util/responder"
	"github.com/munnaMia/nidaa/util/validate"
)

type Handler struct {
	uc        *usecase.UserUseCase
	jshlp     *jsonhelper.JSONHelper
	responder responder.Responder
	validate  validate.Validator
}

// create a user handler
func NewHandler(
	uc *usecase.UserUseCase,
	jshlp *jsonhelper.JSONHelper,
	res responder.Responder,
	valid validate.Validator,
) *Handler {
	return &Handler{
		uc:        uc,
		jshlp:     jshlp,
		responder: res,
		validate:  valid,
	}
}
