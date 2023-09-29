package users

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

const routeName = "users"

func Route(r chi.Router, h *UserHandler, authHandler func(http.Handler) http.Handler) {
	str := fmt.Sprintf("/api/v1/%s", routeName)

	r.Route(str, func(r chi.Router) {
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
