package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

const version = "1.0.6"

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"code\": 0, \"message\": \"Hello World\", \"data\": {\"version\": \"%s\", \"date\": \"%s\"}}", version, time.Now())
}

func loadConfig() {
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	loadConfig()

	http.HandleFunc("/", hello)
	addr := ":8080"
	fmt.Printf("server is listening in %s\n", addr)
	http.ListenAndServe(addr, nil)
}
