package apiserver

import "github.com/sirupsen/logrus"

type ApiServer struct {
	config *Config
	logger *logrus.Logger
}

func NewServer(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return nil
}

func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
