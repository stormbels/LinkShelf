package link

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/stormbels/linkshelf/internal/http-server/handlers/redirect"
	"github.com/stormbels/linkshelf/internal/http-server/viewstate"
	"github.com/stormbels/linkshelf/internal/service"
)

func updateRedirectURL(r *http.Request) string {
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

func NewUpdateHandler(linkService *service.LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawID, id, err := linkIDFromRequest(r)
		if err != nil {
			http.Error(w, "invalid link id", http.StatusBadRequest)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "failed to read form", http.StatusBadRequest)
			return
		}

		form := parseLinkForm(r)

		if err := linkService.UpdateLink(r.Context(), id, form.URL, form.Title, form.Tags); err != nil {
			handleUpdateError(w, r, linkService, err, rawID, form)
			return
		}

		redirectURL := updateRedirectURL(r)

		http.Redirect(w, r, redirect.WithSuccess(redirectURL, viewstate.SuccessUpdated), http.StatusSeeOther)
	}
}

func handleUpdateError(w http.ResponseWriter, r *http.Request, linkService *service.LinkService, err error, rawID string, form linkForm) {
	switch {
	case errors.Is(err, service.ErrLinkNotChanged):
		redirectURL := updateRedirectURL(r)

		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
	case errors.Is(err, service.ErrEmptyURL), errors.Is(err, service.ErrInvalidURL):
		redirectToEditError(w, r, viewstate.ErrorInvalidEdit, rawID, form)
	case errors.Is(err, service.ErrBrokenURL):
		redirectToEditError(w, r, viewstate.ErrorBrokenEdit, rawID, form)
	case errors.Is(err, service.ErrLinkAlreadyExists):
		if !handleDuplicateError(w, r, linkService, form.URL) {
			return
		}
	default:
		http.Error(w, "failed to update link", http.StatusInternalServerError)
	}
}

func redirectToEditError(w http.ResponseWriter, r *http.Request, errorType string, rawID string, form linkForm) {
	redirectToFormError(w, r, errorType, rawID, form)
}
