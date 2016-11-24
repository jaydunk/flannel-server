package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/ifrit/sigmon"
	"github.com/tedsuo/rata"
)

func main() {
	fmt.Println("hello from flannel server")
	routes := rata.Routes{
		{Name: "hello", Method: "GET", Path: "/"},
	}

	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
		return
	})

	handlers := rata.Handlers{
		"hello": helloHandler,
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
