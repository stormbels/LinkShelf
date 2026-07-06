package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/stormbels/linkshelf/internal/http-server/handlers/home"
	linkhandler "github.com/stormbels/linkshelf/internal/http-server/handlers/link"
	"github.com/stormbels/linkshelf/internal/http-server/handlers/notfound"
	"github.com/stormbels/linkshelf/internal/service"
)

func newRouter(log *slog.Logger, linkService *service.LinkService) http.Handler {
	router := chi.NewRouter()

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	router.Get("/", home.New(log, linkService))

	router.Get("/search", linkhandler.NewSearchHandler(linkService))

	router.Post("/links", linkhandler.NewCreateHandler(linkService))

	router.Post("/links/{id}/update", linkhandler.NewUpdateHandler(linkService))

	router.Post("/links/{id}/delete", linkhandler.NewDeleteHandler(linkService))

	router.NotFound(notfound.New(log))

	return router
}
