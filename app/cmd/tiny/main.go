package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kobayashiyabako16g/tiny-go/internal/domain/repository"
	"github.com/kobayashiyabako16g/tiny-go/internal/handler"
	"github.com/kobayashiyabako16g/tiny-go/internal/handler/middleware"
	"github.com/kobayashiyabako16g/tiny-go/pkg/db"
	"github.com/kobayashiyabako16g/tiny-go/pkg/logger"
)

func NewDB(ctx context.Context) (client *db.Client, err error) {
	appEnv := strings.ToLower(os.Getenv("APP_ENV"))
	if appEnv == "local" {
		client, err = db.NewClient("sqlite", "./db/app.db")
	} else if appEnv == "development" {
		client, err = db.NewClient("postgres", "postgres://postgres:postgres@db:5432/app")
	} else {
		err = fmt.Errorf("DB APP_ENV not definision")
	}
	return client, err
}

func main() {
	ctx := context.Background()
	// db setup (ref: makefile)
	client, err := NewDB(ctx)
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
	logger.Info(ctx, fmt.Sprintf("Starting server at http://localhost%s", port))
	// server start
	if err := http.ListenAndServe(port, server); err != nil {
		panic(err)
	}
}
