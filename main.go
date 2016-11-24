package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jaydunk/flannel-server/handlers"

	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/ifrit/sigmon"
	"github.com/tedsuo/rata"
)

const version = "/v1"
const prefix = "/coreos.com/network"

func main() {
	routePrefix := version + prefix
	routes := rata.Routes{
		{Name: "hello", Method: "GET", Path: "/"},
		{Name: "config", Method: "GET", Path: routePrefix + "config"},
		{Name: "acquireLease", Method: "POST", Path: routePrefix + "leases"},
		{Name: "watchLease", Method: "GET", Path: routePrefix + "subnet"},
		{Name: "watchLeases", Method: "GET", Path: routePrefix + "leases"},
	}

	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
		return
	})
	configHandler := &handlers.ConfigHandler{}
	acquireLeaseHandler := &handlers.AcquireLeaseHandler{}
	watchLeaseHandler := &handlers.WatchLeaseHandler{}
	watchLeasesHandler := &handlers.WatchLeasesHandler{}

	handlers := rata.Handlers{
		"hello":        helloHandler,
		"config":       configHandler,
		"acquireLease": acquireLeaseHandler,
		"watchLease":   watchLeaseHandler,
		"watchLeases":  watchLeasesHandler,
	}
	router, err := rata.NewRouter(routes, handlers)
	if err != nil {
		log.Fatalf("unable to create rata Router: %s", err)
	}

	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 5758)
	server := http_server.New(addr, router)

	members := grouper.Members{
		{"http_server", server},
	}

	group := grouper.NewOrdered(os.Interrupt, members)
	monitor := ifrit.Invoke(sigmon.New(group))

	fmt.Println("starting server")
	err = <-monitor.Wait()
	if err != nil {
		log.Fatalf("running server: %s", err)
	}
}
