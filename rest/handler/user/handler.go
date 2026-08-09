package user

import (
	"github.com/munnaMia/nidaa/internal/usecase"
	jsonhelper "github.com/munnaMia/nidaa/util/jsonHelper"
	"github.com/munnaMia/nidaa/util/responder"
	"github.com/munnaMia/nidaa/util/validate"
)

type Handler struct {
	uc        *usecase.UserUseCase
	responder responder.Responder
	jshlp     *jsonhelper.JSONHelper
	validate  validate.Validator
}

// create a user handler
func NewHandler(
	uc *usecase.UserUseCase,
	res responder.Responder,
	jshlp *jsonhelper.JSONHelper,
	valid validate.Validator,
) *Handler {
	return &Handler{
		uc:        uc,
		responder: res,
		jshlp:     jshlp,
		validate:  valid,
	}
}
