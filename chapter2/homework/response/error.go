package response

import (
	"net/http"
)

const errorMessageField = "message"

// WriteErrorString wraps a massage with ErrorsEnvelope and writes to the http.ResponseWriter
func WriteErrorString(w http.ResponseWriter, code int, msg string) error {
	return WriteJSON(w, code, NewErrorsEnvelope().Set(errorMessageField, msg))
}
