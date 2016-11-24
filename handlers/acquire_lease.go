package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/coreos/flannel/subnet"
)

type AcquireLeaseHandler struct{}

func (h *AcquireLeaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, subnet.Lease{})
	return
}

func writeResponse(w http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}
