package response

import "net/http"

// WriteSuccess writes 200 code and payload to the http.ResponseWriter
func WriteSuccess(w http.ResponseWriter, payload interface{}) error {
	return WriteJSON(w, http.StatusOK, payload)
}
