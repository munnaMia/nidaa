package responder

import "net/http"

type Meta struct {
	TotalCount int64 `json:"totalCount"`
	Limit      int   `json:"limit"`
	Offset     int   `json:"offset"`
}

type ValidationErr struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}

type jsonEnvelop struct {
	Data any   `json:"data"`
	Meta *Meta `json:"meta"`
}

type errorEnvelop struct {
	Error struct {
		Code    int             `json:"code"`
		Massage string          `json:"massage"`
		Details []ValidationErr `json:"detials,omitempty"`
	} `json:"error"`
}

type Responder interface {
	SendResponse(w http.ResponseWriter, code int, data any, meta *Meta) // data any type receive a ponter type keep in mind
	SendError(w http.ResponseWriter, code int, msg string, details ...ValidationErr)
}
