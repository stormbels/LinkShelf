package notfound

import (
	"html/template"
	"log/slog"
	"net/http"
)

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r

		tmpl, err := template.ParseFiles("web/templates/not-found.html")
		if err != nil {
			log.Error("failed to parse 404 template", slog.String("error", err.Error()))
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNotFound)

		if err := tmpl.Execute(w, nil); err != nil {
			log.Error("failed to render 404 template", slog.String("error", err.Error()))
		}
	}
}
