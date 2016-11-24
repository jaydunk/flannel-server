package handlers

import (
	"net/http"

	"github.com/coreos/flannel/subnet"
)

type RenewLeaseHandler struct{}

func (h *RenewLeaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, subnet.Lease{})
	return
}
