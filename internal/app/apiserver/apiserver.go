package apiserver

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ApiServer struct {
	config Config
	logger *logrus.Logger
	router chi.Router
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
	s.router.Get("/", s.Root)
}
