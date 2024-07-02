package config

import (
	"net/http"

	"github.com/shironxn/go-clean-arch/internal/domain"
)

type Server struct {
	Host    string
	Port    string
	Handler Handler
}

type Handler struct {
	Author domain.AuthorHandler
}

func NewServer(cfg ServerConfig, handler Handler) *Server {
	return &Server{
		Host:    cfg.Host,
		Port:    cfg.Port,
		Handler: handler,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	router.HandleFunc("/authors", s.Handler.Author.Create)

	server := http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: router,
	}

	return server.ListenAndServe()
}
