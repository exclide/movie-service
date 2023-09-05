package apiserver

import (
	"github.com/exclide/movie-service/internal/app/movies/controller"
	"github.com/exclide/movie-service/internal/app/movies/repository"
	"github.com/exclide/movie-service/internal/app/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type ApiServer struct {
	config Config
	logger *logrus.Logger
	router chi.Router
	store  *store.Store
}

func NewServer(config Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: chi.NewRouter(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *ApiServer) Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root."))
}

func (s *ApiServer) configureRouter() {
	// A good base middleware stack
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.router.Use(middleware.Timeout(60 * time.Second))

	s.router.Get("/", s.Root)

	mr := repository.NewMovieRepository(s.store)
	h := controller.NewHandler(&mr)

	s.router.Route("/api/v1/movies", func(r chi.Router) {
		r.Get("/", h.GetMovies)

		r.Post("/", h.CreateMovie)

		r.Route("/{movieID}", func(r chi.Router) {
			r.Use(h.MovieCtx)
			r.Get("/", h.GetMovie)
			//r.Put("/", h.UpdateMovie)
			r.Delete("/", h.DeleteMovie)
		})
	})
}

func (s *ApiServer) configureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}
