package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kobayashiyabako16g/tiny-go/internal/domain/model"
	"github.com/kobayashiyabako16g/tiny-go/internal/domain/repository"
	"github.com/kobayashiyabako16g/tiny-go/pkg/logger"
)

type User interface {
	GetUser() http.HandlerFunc
	AddUser() http.HandlerFunc
}

type userHandler struct {
	users repository.Users
}

func NewUserHandler(users repository.Users) User {
	return &userHandler{users}
}

type response struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	User    model.User `json:"user"`
}

func (h *userHandler) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			logger.Info(ctx, "Bad Requset", err)
			http.Error(w, "Bad Requset", http.StatusBadRequest)
			return
		}

		user, err := h.users.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, "Repository Error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if user == nil {
			logger.Info(ctx, "Not Found")
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		res := response{
			Status:  "ok",
			Message: "Success Get User",
			User:    *user,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			logger.Error(ctx, "Failed to write response", err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}

func (h *userHandler) AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var user model.User

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			logger.Info(ctx, "Binding Error", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if !user.IsValidEmail() {
			logger.Info(ctx, "Validation Error")
			http.Error(w, "Bad Request (email)", http.StatusBadRequest)
			return
		}

		// create User
		err := h.users.Create(ctx, &user)
		if err != nil {
			logger.Info(ctx, "User Create Error")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return

		}

		res := response{
			Status:  "ok",
			Message: "Success Create User",
			User:    user,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			logger.Error(ctx, "Failed to write response", err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}
