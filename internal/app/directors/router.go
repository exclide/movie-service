package directors

import (
	"fmt"
	"github.com/go-chi/chi"
)

const routeName = "directors"

func Route(r chi.Router, h *DirectorHandler) {
	str := fmt.Sprintf("/api/v1/%s", routeName)

	r.Route(str, func(r chi.Router) {
		r.Get("/", h.GetDirectors)

		r.Post("/", h.CreateDirector)

		r.Route("/{dirID}", func(r chi.Router) {
			r.Use(h.DirectorCtx)
			r.Get("/", h.GetDirector)
			//r.Put("/", dirHandler.UpdateMovie)
			r.Delete("/", h.DeleteDirector)
		})
	})
}
