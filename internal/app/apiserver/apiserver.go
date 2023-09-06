package apiserver

import (
	controller2 "github.com/exclide/movie-service/internal/app/directors/controller"
	repository2 "github.com/exclide/movie-service/internal/app/directors/repository"
	"github.com/exclide/movie-service/internal/app/movies/controller"
	"github.com/exclide/movie-service/internal/app/movies/repository"
	"github.com/exclide/movie-service/internal/app/store"
	controller3 "github.com/exclide/movie-service/internal/app/users/controller"
	repository3 "github.com/exclide/movie-service/internal/app/users/repository"
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
	w.Write([]byte("hello world"))
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

	movieRepo := repository.NewMovieRepository(s.store)
	movieHandler := controller.NewMovieHandler(&movieRepo)
	dirRepo := repository2.NewDirectorRepository(s.store)
	dirHandler := controller2.NewDirectorHandler(&dirRepo)
	userRepo := repository3.NewUserRepository(s.store)
	userHandler := controller3.NewUserHandler(&userRepo)

	s.router.Route("/api/v1/movies", func(r chi.Router) {
		r.Get("/", movieHandler.GetMovies)

		r.Post("/", movieHandler.CreateMovie)

		r.Route("/{movieID}", func(r chi.Router) {
			r.Use(movieHandler.MovieCtx)
			r.Get("/", movieHandler.GetMovie)
			//r.Put("/", movieHandler.UpdateMovie)
			r.Delete("/", movieHandler.DeleteMovie)
		})
	})

	s.router.Route("/api/v1/directors", func(r chi.Router) {
		r.Get("/", dirHandler.GetDirectors)

		r.Post("/", dirHandler.CreateDirector)

		r.Route("/{dirID}", func(r chi.Router) {
			r.Use(dirHandler.DirectorCtx)
			r.Get("/", dirHandler.GetDirector)
			//r.Put("/", dirHandler.UpdateMovie)
			r.Delete("/", dirHandler.DeleteDirector)
		})
	})

	s.router.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)

		r.Post("/", userHandler.CreateUser)

		r.Route("/{dirID}", func(r chi.Router) {
			r.Use(userHandler.UserCtx)
			r.Get("/", userHandler.GetUser)
			//r.Put("/", userHandler.UpdateMovie)
			r.Delete("/", userHandler.DeleteUser)
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
