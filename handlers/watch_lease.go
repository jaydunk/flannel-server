package handlers

import (
	"net/http"

	"github.com/coreos/flannel/subnet"
)

type WatchLeaseHandler struct{}

func (h *WatchLeaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, subnet.LeaseWatchResult{})
	return
}
