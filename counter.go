package main

import (
	"fmt"
	"net/http"
)

var number int = 0

func get(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "%d\n", number)
}

func incr(response http.ResponseWriter, request *http.Request) {
	number = number + 1;
	fmt.Fprintf(response, "%d\n", number)
}

func reset(response http.ResponseWriter, request *http.Request) {
	number = 0;
	fmt.Fprintf(response, "%d\n", number)
}

func main() {
	http.HandleFunc("/",      get)
	http.HandleFunc("/get",   get)
	http.HandleFunc("/incr",  incr)
	http.HandleFunc("/reset", reset)
	http.ListenAndServe(":8080", nil)
}