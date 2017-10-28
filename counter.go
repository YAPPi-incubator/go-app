package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

var (
	countValue    countDriver = &inMemoryCount{0}
	listenAddress             = flag.String("listen-address", ":8080", "Address on which to expose service.")
)

type countDriver interface {
	get() int64
	inc() int64
	reset() int64
}

type inMemoryCount struct {
	value int64
}

func (c inMemoryCount) get() int64 {
	return c.value
}

func (c *inMemoryCount) inc() int64 {
	c.value++

	return c.value
}

func (c *inMemoryCount) reset() int64 {
	c.value = 0

	return c.value
}

type serviceHandler struct {
}

func (h serviceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var value int64

	switch req.URL.Path {
	case "/inc":
		value = countValue.inc()
	case "/reset":
		value = countValue.reset()
	default:
		value = countValue.get()
	}

	fmt.Fprintf(w, "%d\n", value)
}

func main() {
	flag.Parse()

	log.Infof("binding to %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, serviceHandler{}))
}
