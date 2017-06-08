package handler

import (
	"gochapter/chapter2/homework/response"
	"log"
	"net/http"
)

type recoveryHandler struct {
	handler http.Handler
}

// NewRecovery return recovery handler
func NewRecovery(h http.Handler) http.Handler {
	return &recoveryHandler{h}
}

func (h *recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// recover allows you to continue execution in case of panic
		if err := recover(); err != nil {
			response.WriteErrorString(w, http.StatusInternalServerError, "Internal Server Error")
			log.Printf("PANIC! %s %s: %v\n", r.Method, r.RequestURI, err)
		}
	}()

	h.handler.ServeHTTP(w, r)
}
