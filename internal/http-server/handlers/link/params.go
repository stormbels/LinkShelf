package link

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func linkIDFromRequest(r *http.Request) (rawID string, id int64, err error) {
	rawID = chi.URLParam(r, "id")
	id, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return "", 0, err
	}

	return rawID, id, nil
}
