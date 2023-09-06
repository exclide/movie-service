package controller

import (
	"context"
	"encoding/json"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/movies"
	"github.com/exclide/movie-service/internal/app/utils"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	repository movies.Repository
}

func NewMovieHandler(r movies.Repository) *MovieHandler {
	return &MovieHandler{r}
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("movie").(*model.Movie)

	err := json.NewEncoder(w).Encode(mv)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
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

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var mv model.Movie

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

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("movie").(*model.Movie)

	err := h.repository.DeleteById(r.Context(), mv.Id)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	json.NewEncoder(w).Encode("Delete ok")
}

func (h *MovieHandler) MovieCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		movieID, _ := strconv.Atoi(chi.URLParam(r, "movieID"))
		movie, err := h.repository.GetById(r.Context(), movieID)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		ctx := context.WithValue(r.Context(), "movie", movie)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
