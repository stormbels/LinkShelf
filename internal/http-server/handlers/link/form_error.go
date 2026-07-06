package link

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/stormbels/linkshelf/internal/http-server/handlers/redirect"
)

func redirectToFormError(w http.ResponseWriter, r *http.Request, errorType string, rawID string, form linkForm) {
	query := url.Values{}
	query.Set("error", errorType)

	if rawID != "" {
		query.Set("id", rawID)
	}

	query.Set("url", strings.TrimSpace(form.URL))
	query.Set("title", form.Title)
	query.Set("tags", form.Tags)

	redirect.ToTargetOrHome(w, r, r.URL.Query().Get("redirect"), query)
}
