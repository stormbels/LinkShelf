package link

import (
	"errors"
	"net/http"

	"github.com/stormbels/linkshelf/internal/http-server/handlers/redirect"
	"github.com/stormbels/linkshelf/internal/http-server/viewstate"
	"github.com/stormbels/linkshelf/internal/service"
)

func NewCreateHandler(linkService *service.LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "failed to read form", http.StatusBadRequest)
			return
		}

		form := parseLinkForm(r)

		if err := linkService.CreateLink(r.Context(), form.URL, form.Title, form.Tags); err != nil {
			handleCreateError(w, r, linkService, err, form)
			return
		}

		redirect.ToHomeWithSuccess(w, r, viewstate.SuccessCreated)
	}
}

func handleCreateError(w http.ResponseWriter, r *http.Request, linkService *service.LinkService, err error, form linkForm) {
	switch {
	case errors.Is(err, service.ErrEmptyURL), errors.Is(err, service.ErrInvalidURL):
		redirectToCreateError(w, r, viewstate.ErrorInvalidCreate, form)
	case errors.Is(err, service.ErrBrokenURL):
		redirectToCreateError(w, r, viewstate.ErrorBrokenCreate, form)
	case errors.Is(err, service.ErrLinkAlreadyExists):
		if !handleDuplicateError(w, r, linkService, form.URL) {
			return
		}
	default:
		http.Error(w, "failed to save link", http.StatusInternalServerError)
	}
}

func redirectToCreateError(w http.ResponseWriter, r *http.Request, errorType string, form linkForm) {
	redirectToFormError(w, r, errorType, "", form)
}
