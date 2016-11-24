package handlers

import (
	"net/http"

	"github.com/coreos/flannel/subnet"
)

type WatchLeasesHandler struct{}

func (h *WatchLeasesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// always write the current state to the lease watch result snapshot
	lwResult := subnet.LeaseWatchResult{
		Snapshot: []subnet.Lease{},
		Cursor:   "1",
	}
	writeResponse(w, lwResult)
	return
}
