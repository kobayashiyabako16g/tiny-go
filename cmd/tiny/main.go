package main

import (
	"fmt"
	"net/http"

	"github.com/kobayashiyabako16g/tiny-go/internal/domain/repository"
	"github.com/kobayashiyabako16g/tiny-go/internal/handler"
	"github.com/kobayashiyabako16g/tiny-go/internal/handler/middleware"
	"github.com/kobayashiyabako16g/tiny-go/pkg/db"
	"github.com/kobayashiyabako16g/tiny-go/pkg/log"
)

func main() {
	// db setup (ref: makefile)
	client, err := db.NewClient("sqlite", "./db/app.db")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// dependency injection
	userRepository := repository.NewUsersRepository(client)
	userHandler := handler.NewUserHandler(userRepository)

	router := handler.NewRouter(userHandler)
	mux := http.NewServeMux()
	router.HandleRequest(mux)

	server := middleware.LogMiddleware(mux)
	// server setup
	port := ":8080"
	log.Logger.Info(fmt.Sprintf("Starting server at http://localhost%s", port))
	// server start
	if err := http.ListenAndServe(port, server); err != nil {
		panic(err)
	}
}
