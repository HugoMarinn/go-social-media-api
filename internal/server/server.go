package server

import "github.com/HugoMarinn/go-social-media-api/internal/config"

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{cfg}
}

func (s *Server) Run() {

}
