package main

import (
	"fmt"
	"net/http"
	"time"
)

const version = "1.0.1"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"code\": 0, \"message\": \"Hello World\", \"data\": {\"version\": \"%s\", \"date\": \"%s\"}}", version, time.Now())
}

func main() {
	http.HandleFunc("/", hello)
	addr := ":8080"
	fmt.Printf("server is listening in %s\n", addr)
	http.ListenAndServe(addr, nil)
}
