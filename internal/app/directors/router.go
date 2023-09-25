package directors

import "github.com/go-chi/chi"

func Route(r chi.Router, h *DirectorHandler) {
	r.Route("/api/v1/directors", func(r chi.Router) {
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
