package handlers

import "net/http"

type RevokeLeaseHandler struct{}

func (h *RevokeLeaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
	w.WriteHeader(http.StatusOK)
	return
}
