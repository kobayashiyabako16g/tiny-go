package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World")
}

func main() {
	// server setup
	http.HandleFunc("/", handler)
	port := ":8080"
	fmt.Printf("Starting server at http://localhost%s\n", port)

	// server start
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
