package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/stormbels/linkshelf/internal/config"
	"github.com/stormbels/linkshelf/internal/service"
)

func main() {
	log := newLogger()

	cfg, err := config.Load("config/local.yaml")
	if err != nil {
		log.Error("failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	linkStorage, closeStorage, err := newLinkStorage(cfg)
	if err != nil {
		log.Error("failed to create link storage", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer closeStorage()

	linkService := service.NewLinkService(linkStorage)

	router := newRouter(log, linkService)

	server := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	runServer(log, server)
}
