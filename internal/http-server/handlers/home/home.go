package home

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/stormbels/linkshelf/internal/domain"
	"github.com/stormbels/linkshelf/internal/http-server/viewstate"
	"github.com/stormbels/linkshelf/internal/service"
)

type PageData struct {
	Links       []domain.Link
	ErrorType   string
	SuccessType string
}

func New(log *slog.Logger, linkService *service.LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		links, err := linkService.ListLinks(r.Context())
		if err != nil {
			log.Error("failed to list links", slog.String("error", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles(
			"web/templates/home.html",
			"web/templates/partials/home-scripts.html",
			"web/templates/partials/link-modals.html",
			"web/templates/partials/delete-confirm-dialog.html",
			"web/templates/partials/state-dialogs.html",
			"web/templates/partials/shelf-status.html",
			"web/templates/partials/meme-section.html",
			"web/templates/partials/links-section.html",
		)
		if err != nil {
			log.Error("failed to parse home template", slog.String("error", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Links:       links,
			ErrorType:   viewstate.ErrorType(r.URL.Query().Get("error")),
			SuccessType: viewstate.SuccessType(r.URL.Query().Get("success")),
		}

		if err := tmpl.Execute(w, data); err != nil {
			log.Error("failed to render home template", slog.String("error", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
