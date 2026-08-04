package responder // send http response to a response writer.

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
)

type httpResponder struct {
}

func NewHttpResponder() Responder {
	return &httpResponder{}
}

func (htres *httpResponder) SendResponse(w http.ResponseWriter, code int, data any, meta *Meta) {
	w.Header().Set("Content-Type", "application/json")

	//Encode into a buffer first (does NOT write to network yet)
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(jsonEnvelop{
		Data: data,
		Meta: meta,
	}); err != nil {
		slog.Error("Could not sending JSON response", "Err", err)
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	_, _ = w.Write(buf.Bytes())
}

// send http error response to a writer
func (htres *httpResponder) SendError(w http.ResponseWriter, code int, msg string, details ...ValidationErr) {
	w.Header().Set("Content-Type", "application/json")

	var buf bytes.Buffer

	errEnv := errorEnvelop{}

	errEnv.Error.Code = code
	errEnv.Error.Massage = msg
	errEnv.Error.Details = details

	if err := json.NewEncoder(&buf).Encode(errEnv); err != nil {
		slog.Error("Could not sending JSON response", "Err", err)
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	_, _ = w.Write(buf.Bytes())
}
