package response

import (
	"encoding/json"
	"net/http"
)

// writeJSON writes response code and json encoded payload to the http.ResponseWriter
func WriteJSON(w http.ResponseWriter, code int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if nil != payload {
		encoder := json.NewEncoder(w)
		return encoder.Encode(payload)
	}

	return nil
}
