package controller

import (
	"context"
	"encoding/json"
	"github.com/exclide/movie-service/internal/app/directors"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type DirectorHandler struct {
	repository directors.Repository
}

func NewDirectorHandler(r directors.Repository) *DirectorHandler {
	return &DirectorHandler{r}
}

func (h *DirectorHandler) GetDirector(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("dir").(*model.Director)

	err := json.NewEncoder(w).Encode(mv)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *DirectorHandler) GetDirectors(w http.ResponseWriter, r *http.Request) {
	mv, err := h.repository.GetAll(r.Context())

	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(mv)

	if err != nil {
		log.Fatal(err)
	}
}

func (h *DirectorHandler) CreateDirector(w http.ResponseWriter, r *http.Request) {
	var mv model.Director

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

func (h *DirectorHandler) DeleteDirector(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("dir").(*model.Director)

	err := h.repository.DeleteById(r.Context(), mv.Id)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Delete ok")
}

func (h *DirectorHandler) DirectorCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dirID, _ := strconv.Atoi(chi.URLParam(r, "dirID"))
		dir, err := h.repository.GetById(r.Context(), dirID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "dir", dir)
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
