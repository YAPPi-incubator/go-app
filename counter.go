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

func get(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%d\n", number.get())
}

func inc(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%d\n", number.inc())
}

func reset(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%d\n", number.reset())
}

func main() {
	flag.Parse()

	log.Infof("binding to %s", *listenAddress)
	http.HandleFunc("/", get)
	http.HandleFunc("/get", get)
	http.HandleFunc("/inc", inc)
	http.HandleFunc("/reset", reset)

	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
