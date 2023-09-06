package controller

import (
	"context"
	"encoding/json"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/users"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type UserHandler struct {
	repository users.Repository
}

func NewUserHandler(r users.Repository) *UserHandler {
	return &UserHandler{r}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("user").(*model.User)

	err := json.NewEncoder(w).Encode(mv)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	mv, err := h.repository.GetAll(r.Context())

	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(mv)

	if err != nil {
		log.Fatal(err)
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var mv model.User

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	create, err := h.repository.Create(r.Context(), &mv)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(create)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("user").(*model.User)

	err := h.repository.DeleteByLogin(r.Context(), mv.Login)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Delete ok")
}

func (h *UserHandler) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		user, err := h.repository.GetByLogin(r.Context(), userID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		//w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
