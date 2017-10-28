package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

var (
	number        = countValue{0}
	listenAddress = flag.String("listen-address", ":8080", "Address on which to expose service.")
)

type countValue struct {
	value int64
}

func (c countValue) get() int64 {
	return c.value
}

func (c *countValue) inc() int64 {
	c.value++

	return c.value
}

func (c *countValue) reset() int64 {
	c.value = 0

	return c.value
}

type serviceHandler struct {
}

func (h serviceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/inc":
		fmt.Fprintf(w, "%d\n", number.inc())
	case "/reset":
		fmt.Fprintf(w, "%d\n", number.reset())
	default:
		fmt.Fprintf(w, "%d\n", number.get())
	}
}

func main() {
	flag.Parse()

	log.Infof("binding to %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, serviceHandler{}))
}
