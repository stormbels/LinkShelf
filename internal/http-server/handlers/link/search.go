package link

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/stormbels/linkshelf/internal/domain"
	"github.com/stormbels/linkshelf/internal/http-server/viewstate"
	"github.com/stormbels/linkshelf/internal/service"
)

type SearchPageData struct {
	Links       []domain.Link
	Query       string
	IsSelected  bool
	ErrorType   string
	SuccessType string
}

func NewSearchHandler(linkService *service.LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")

		links, err := linkService.SearchLinks(r.Context(), query)
		if err != nil {
			http.Error(w, "failed to search links", http.StatusInternalServerError)
			return
		}

		isSelected := false
		selectedID, err := strconv.ParseInt(r.URL.Query().Get("selected"), 10, 64)
		if err == nil {
			links = filterLinksByID(links, selectedID)
			isSelected = true
		}

		errorType := viewstate.ErrorType(r.URL.Query().Get("error"))
		successType := viewstate.SuccessType(r.URL.Query().Get("success"))

		tmpl, err := template.ParseFiles(
			"web/templates/search.html",
			"web/templates/partials/search-results-section.html",
			"web/templates/partials/search-edit-modal.html",
			"web/templates/partials/search-delete-confirm-dialog.html",
			"web/templates/partials/search-error-dialog.html",
			"web/templates/partials/search-success-dialog.html",
		)
		if err != nil {
			http.Error(w, "failed to parse template", http.StatusInternalServerError)
			return
		}

		data := SearchPageData{
			Links:       links,
			Query:       query,
			IsSelected:  isSelected,
			ErrorType:   errorType,
			SuccessType: successType,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	}
}

func filterLinksByID(links []domain.Link, selectedID int64) []domain.Link {
	for _, link := range links {
		if link.ID == selectedID {
			return []domain.Link{link}
		}
	}

	return []domain.Link{}
}
