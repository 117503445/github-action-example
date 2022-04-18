package src

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

const Version = "1.0.10"

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{\"code\": 0, \"message\": \"Hello World\", \"data\": {\"version\": \"%s\", \"date\": \"%s\"}}", Version, time.Now())
}

func LoadConfig() {
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func Start() {
	LoadConfig()

	http.HandleFunc("/", Hello)
	addr := ":8080"
	fmt.Printf("server is listening in %s\n", addr)
	http.ListenAndServe(addr, nil)
}
