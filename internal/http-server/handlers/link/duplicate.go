package link

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/stormbels/linkshelf/internal/http-server/handlers/redirect"
	"github.com/stormbels/linkshelf/internal/http-server/viewstate"
	"github.com/stormbels/linkshelf/internal/service"
)

func handleDuplicateError(w http.ResponseWriter, r *http.Request, linkService *service.LinkService, rawURL string) bool {
	cleanURL := strings.TrimSpace(rawURL)

	query := url.Values{}
	query.Set("error", viewstate.ErrorDuplicate)
	query.Set("url", cleanURL)

	existingLink, found, err := linkService.FindDuplicateByRawURL(r.Context(), cleanURL)
	if err != nil {
		http.Error(w, "failed to find duplicate link", http.StatusInternalServerError)
		return false
	}

	if found {
		linkID := strconv.FormatInt(existingLink.ID, 10)
		tags := strings.Join(existingLink.Tags, ", ")

		query.Set("id", linkID)
		query.Set("url", existingLink.URL)
		query.Set("title", existingLink.Title)
		query.Set("tags", tags)
		query.Set("existing_id", linkID)
		query.Set("existing_url", existingLink.URL)
		query.Set("existing_title", existingLink.Title)
		query.Set("existing_tags", tags)
	}

	redirect.ToTargetOrHome(w, r, r.URL.Query().Get("redirect"), query)
	return true
}
