package handlers

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/coreos/flannel/pkg/ip"
	"github.com/coreos/flannel/subnet"
)

type AcquireLeaseHandler struct{}

func (h *AcquireLeaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	attrs := subnet.LeaseAttrs{}
	err := json.NewDecoder(r.Body).Decode(&attrs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, net, err := net.ParseCIDR("10.255.99.0/24")
	if err != nil {
		panic(err)
	}

	lease := subnet.Lease{
		Subnet:     ip.FromIPNet(net),
		Attrs:      attrs,
		Expiration: time.Now().Add(24 * time.Hour),
	}
	writeResponse(w, lease)
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
