package http

import (
	"github.com/go-chi/chi"
)

type Server struct {
	router *chi.Mux
}

func (s *Server) Router() *chi.Mux {
	return s.router
}

func NewServer(router *chi.Mux) *Server {
	return &Server{router: router}
}
