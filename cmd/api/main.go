package main

import (
	"log"

	"github.com/HugoMarinn/go-social-media-api/internal/config"
	"github.com/HugoMarinn/go-social-media-api/internal/server"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("fail to load config: ", err)
	}

	srv := server.New(cfg)
	srv.Run()
}
