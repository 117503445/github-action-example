package main

import (
	"fmt"
	"net/http"
)

const version = "1.0.0"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"code\": 0, \"message\": \"Hello World\", \"data\": {\"version\": \"%s\"}}", version)
}

func main() {
	http.HandleFunc("/", hello)
	addr := ":8080"
	fmt.Printf("http server listening in %s\n", addr)
	http.ListenAndServe(addr, nil)
}
