package main

import (
	"encoding/json"
	"fmt"
	"github.com/kobayashiyabako16g/tiny-go/pkg/log"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	res := Response{
		Status:  "ok",
		Message: "Hello, World",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Logger.Error("Failed to write response", "error", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
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
