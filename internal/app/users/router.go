package users

import (
	"github.com/go-chi/chi"
	"net/http"
)

func Route(r chi.Router, h *UserHandler, authHandler func(http.Handler) http.Handler) {
	r.Route("/api/v1/directors", func(r chi.Router) {
		r.Get("/", h.GetUsers)

		r.With(authHandler).Post("/", h.CreateUser)

		r.Route("/{dirID}", func(r chi.Router) {
			r.Use(h.UserCtx)
			r.Get("/", h.GetUser)
			//r.Put("/", dirHandler.UpdateMovie)
			r.Delete("/", h.DeleteUser)
		})
	})
}
