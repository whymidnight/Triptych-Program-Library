package structs

import (
	"encoding/json"
	"net/http"
)

type ResponseT struct {
	Message interface{} `json:"message"`
}

type Response interface {
	WriteNewResponse(w *http.ResponseWriter, body interface{}, status int)
}

func (r *ResponseT) WriteNewResponse(w *http.ResponseWriter, body interface{}, status int) {
	r.Message = body

	responseJson, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		(*w).WriteHeader(501)
		return
	}

	(*w).WriteHeader(status)
	(*w).Write(responseJson)
}
