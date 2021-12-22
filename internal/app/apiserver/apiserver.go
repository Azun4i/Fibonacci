package apiserver

import "net/http"

func Start(config *Config) error {
	s := newServer()
	s.logger.Info("starting server")
	return http.ListenAndServe(config.BindAddr, s)
}
