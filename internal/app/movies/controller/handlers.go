package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/movies"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	repository movies.Repository
}

func NewHandler(r movies.Repository) *MovieHandler {
	return &MovieHandler{r}
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("movie").(*model.Movie)

	err := json.NewEncoder(w).Encode(mv)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GAVNO")

	mv, err := h.repository.GetAll(r.Context())
	fmt.Println(mv)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(mv)

	if err != nil {
		log.Fatal(err)
	}
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var mv model.Movie

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {

}

func (h *MovieHandler) MovieCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		movieID, _ := strconv.Atoi(chi.URLParam(r, "movieID"))
		movie, err := h.repository.GetById(r.Context(), movieID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "movie", movie)
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
