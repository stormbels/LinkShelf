package link

import (
	"net/http"

	"github.com/stormbels/linkshelf/internal/http-server/handlers/redirect"
	"github.com/stormbels/linkshelf/internal/http-server/viewstate"
	"github.com/stormbels/linkshelf/internal/service"
)

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

		redirect.ToHomeWithSuccess(w, r, viewstate.SuccessDeleted)
	}
}
