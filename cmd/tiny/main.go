package main

import (
	"encoding/json"
	"fmt"
	"github.com/kobayashiyabako16g/tiny-go/pkg/db"
	"github.com/kobayashiyabako16g/tiny-go/pkg/log"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	User    User   `json:"user"`
}

type User struct {
	Id    int
	Name  string
	Email string
}

func handler(w http.ResponseWriter, r *http.Request, dbClient *db.Client) {
	stmt, err := dbClient.Prepare("SELECT id, name, email FROM users WHERE id = ?")
	if err != nil {
		log.Logger.Error("Database Error", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	id := r.PathValue("id")
	var user User
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		log.Logger.Error("Database Error", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res := Response{
		Status:  "ok",
		Message: "Hello, World",
		User:    user,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Logger.Error("Failed to write response", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	// db setup (ref: makefile)
	client, err := db.NewClient("sqlite", "./db/app.db")
	if err != nil {
		panic(err)
	}

	// server setup
	http.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, client)
	})
	port := ":8080"
	log.Logger.Info(fmt.Sprintf("Starting server at http://localhost%s\n", port))

	// server start
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
