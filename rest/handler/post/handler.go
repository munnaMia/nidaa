package post

import (
	"github.com/munnaMia/nidaa/internal/usecase"
	jsonhelper "github.com/munnaMia/nidaa/util/jsonHelper"
	"github.com/munnaMia/nidaa/util/responder"
)

type Handler struct {
	uc    *usecase.PostUseCase
	jshlp *jsonhelper.JSONHelper
	res   responder.Responder
}

// create a user handler
func NewHandler(
	uc *usecase.PostUseCase,
	jshlp *jsonhelper.JSONHelper,
	res responder.Responder,
) *Handler {
	return &Handler{
		uc:    uc,
		jshlp: jshlp,
		res:   res,
	}
}
