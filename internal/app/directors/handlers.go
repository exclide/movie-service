package directors

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

type DirectorHandler struct {
	serv Service
}

func NewDirectorHandler(r Service) *DirectorHandler {
	return &DirectorHandler{r}
}

func (h *DirectorHandler) GetDirector(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value(key).(*model.Director)

	err := json.NewEncoder(w).Encode(mv)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *DirectorHandler) GetDirectors(w http.ResponseWriter, r *http.Request) {
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

func (h *DirectorHandler) CreateDirector(w http.ResponseWriter, r *http.Request) {
	var mv model.Director

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

func (h *DirectorHandler) DeleteDirector(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value(key).(*model.Director)

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

const key contextKey = "dir"

func (h *DirectorHandler) DirectorCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dirID, _ := strconv.Atoi(chi.URLParam(r, "dirID"))
		dir, err := h.serv.GetById(r.Context(), dirID)
		if err != nil {
			httpformat.Error(w, r, http.StatusBadRequest, err)
			return
		}
		ctx := context.WithValue(r.Context(), key, dir)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
