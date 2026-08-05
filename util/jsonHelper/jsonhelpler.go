package jsonhelper

import (
	"encoding/json"
	"io"
)

type JSONHelper struct {
}

func NewJsonHelper() *JSONHelper {
	return &JSONHelper{}
}

// takes an io.reader and v any type value and assing the json payload form the request to the given v.
func (jh *JSONHelper) Decoder(r io.Reader, v any) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}
