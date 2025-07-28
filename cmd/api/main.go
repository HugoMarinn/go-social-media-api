package main

import (
	"log"

	authHttp "github.com/HugoMarinn/go-social-media-api/internal/auth/delivery/http"
	authRepo "github.com/HugoMarinn/go-social-media-api/internal/auth/repository"
	authUseCase "github.com/HugoMarinn/go-social-media-api/internal/auth/usecase"
	"github.com/HugoMarinn/go-social-media-api/internal/config"
	"github.com/HugoMarinn/go-social-media-api/internal/server"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("fail to load config: ", err)
	}

	srv := server.New(cfg)

	authRepo := authRepo.NewPostgresAuthRepository(cfg.DB)
	authUseCase := authUseCase.NewAuthUseCase(authRepo)
	authHandler := authHttp.NewAuthHandler(authUseCase)

	srv.Run(
		authHandler,
	)
}
