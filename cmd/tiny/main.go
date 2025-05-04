package main

import (
	"fmt"
	"github.com/kobayashiyabako16g/tiny-go/pkg/log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("Hello, World")
}

func main() {
	// server setup
	http.HandleFunc("/", handler)
	port := ":8080"
	log.Logger.Info(fmt.Sprintf("Starting server at http://localhost%s\n", port))

	// server start
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
