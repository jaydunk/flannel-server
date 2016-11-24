package handlers

import (
	"net/http"

	"github.com/coreos/flannel/subnet"
)

type ConfigHandler struct{}

func (h *ConfigHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, subnet.Config{})
	return
}
