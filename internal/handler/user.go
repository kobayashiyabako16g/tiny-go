package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kobayashiyabako16g/tiny-go/internal/domain/model"
	"github.com/kobayashiyabako16g/tiny-go/internal/domain/repository"
	"github.com/kobayashiyabako16g/tiny-go/pkg/log"
)

type User interface {
	GetUser() http.HandlerFunc
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

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Logger.Info("Bad Requset", "info", err)
			http.Error(w, "Bad Requset", http.StatusBadRequest)
			return
		}

		user, err := h.users.FindById(ctx, id)
		if err != nil {
			log.Logger.Error("Repository Error", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if user == nil {
			log.Logger.Info("Not Found")
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
			log.Logger.Error("Failed to write response", "error", err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}
