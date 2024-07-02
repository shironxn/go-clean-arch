package config

import (
	"net/http"
)

type Server struct {
	Host string
	Port string
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		Host: cfg.Host,
		Port: cfg.Port,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	server := http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: router,
	}

	return server.ListenAndServe()
}
