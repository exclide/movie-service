package movies

import (
	"context"
	"encoding/json"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/pkg/httpformat"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	serv Service
}

func NewHandler(r Service) *MovieHandler {
	return &MovieHandler{r}
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value(key).(*model.Movie)

	err := json.NewEncoder(w).Encode(mv)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	mv, err := h.serv.GetAll(r.Context())

	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode(mv)

	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var mv model.Movie

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	create, err := h.serv.Create(r.Context(), &mv)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode(create)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value(key).(*model.Movie)

	err := h.serv.DeleteById(r.Context(), mv.Id)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode("Delete ok")
	if err != nil {
		logrus.Warn(err)
	}
}

type contextKey string

const key contextKey = "movie"

func (h *MovieHandler) MovieCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		movieID, _ := strconv.Atoi(chi.URLParam(r, "movieID"))
		movie, err := h.serv.GetById(r.Context(), movieID)
		if err != nil {
			httpformat.Error(w, r, http.StatusBadRequest, err)
			return
		}
		ctx := context.WithValue(r.Context(), key, movie)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
