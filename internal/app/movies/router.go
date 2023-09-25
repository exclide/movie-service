package movies

import "github.com/go-chi/chi"

func Route(r chi.Router, h *MovieHandler) {
	r.Route("/api/v1/directors", func(r chi.Router) {
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
