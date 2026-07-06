package link

import (
	"net/http"
	"net/url"

	"github.com/stormbels/linkshelf/internal/http-server/handlers/redirect"
	"github.com/stormbels/linkshelf/internal/http-server/viewstate"
	"github.com/stormbels/linkshelf/internal/service"
)

func deleteRedirectURL(r *http.Request) string {
	redirectURL := r.FormValue("redirect")
	if redirectURL != "" {
		return redirectURL
	}

	redirectURL = r.FormValue("target")
	if redirectURL != "" {
		return redirectURL
	}

	refererURL, err := url.Parse(r.Referer())
	if err == nil && refererURL.Host == r.Host {
		return refererURL.RequestURI()
	}

	return "/"
}

func NewDeleteHandler(linkService *service.LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, id, err := linkIDFromRequest(r)
		if err != nil {
			http.Error(w, "invalid link id", http.StatusBadRequest)
			return
		}

		if err := linkService.DeleteLink(r.Context(), id); err != nil {
			http.Error(w, "failed to delete link", http.StatusInternalServerError)
			return
		}

		redirectURL := deleteRedirectURL(r)
		http.Redirect(w, r, redirect.WithSuccess(redirectURL, viewstate.SuccessDeleted), http.StatusSeeOther)
	}
}
