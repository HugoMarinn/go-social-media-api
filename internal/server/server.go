package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HugoMarinn/go-social-media-api/internal/config"
)

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{cfg}
}

func (s *Server) Run() {
	handler := MapRoutes()
	server := http.Server{
		Addr:    ":" + s.cfg.Port,
		Handler: handler,
	}

	log.Printf("server is running in http://localhost:%s/api/v1", s.cfg.Port)
	// Run Server in another goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error initializing server: %v", err)
		}
	}()

	// Waiting Ctrl + C or another Interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("stoping the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to stop the server: %v", err)
	}

	if err := s.cfg.DB.Close(); err != nil {
		log.Fatalf("failed to close db connection: %v", err)
	}
	log.Println("db connection closed successfully")

	log.Println("server finished successfully")
}
