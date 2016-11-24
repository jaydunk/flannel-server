package handlers

import (
	"net/http"

	"github.com/coreos/flannel/subnet"
)

type WatchLeasesHandler struct{}

func (h *WatchLeasesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, subnet.LeaseWatchResult{})
	return
}
