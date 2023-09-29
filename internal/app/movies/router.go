package movies

import (
	"fmt"
	"github.com/go-chi/chi"
)

const routeName = "movies"

func Route(r chi.Router, h *MovieHandler) {
	str := fmt.Sprintf("/api/v1/%s", routeName)

	r.Route(str, func(r chi.Router) {
		r.Get("/", h.GetMovies)

		r.Post("/", h.CreateMovie)

		r.Route("/{dirID}", func(r chi.Router) {
			r.Use(h.MovieCtx)
			r.Get("/", h.GetMovie)
			//r.Put("/", dirHandler.UpdateMovie)
			r.Delete("/", h.DeleteMovie)
		})
	})
}
