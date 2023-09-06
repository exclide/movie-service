package controller

import (
	"context"
	"encoding/json"
	"github.com/exclide/movie-service/internal/app/directors"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/utils"
	"github.com/go-chi/chi"
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
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *DirectorHandler) GetDirectors(w http.ResponseWriter, r *http.Request) {
	mv, err := h.repository.GetAll(r.Context())

	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode(mv)

	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *DirectorHandler) CreateDirector(w http.ResponseWriter, r *http.Request) {
	var mv model.Director

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	create, err := h.repository.Create(r.Context(), &mv)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode(create)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *DirectorHandler) DeleteDirector(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("dir").(*model.Director)

	err := h.repository.DeleteById(r.Context(), mv.Id)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	json.NewEncoder(w).Encode("Delete ok")
}

func (h *DirectorHandler) DirectorCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dirID, _ := strconv.Atoi(chi.URLParam(r, "dirID"))
		dir, err := h.repository.GetById(r.Context(), dirID)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		ctx := context.WithValue(r.Context(), "dir", dir)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
