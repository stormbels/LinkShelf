package redirect

import (
	"net/http"
	"net/url"
)

func ToHome(w http.ResponseWriter, r *http.Request, query url.Values) {
	path := "/"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	http.Redirect(w, r, path, http.StatusSeeOther)
}

func ToHomeWithSuccess(w http.ResponseWriter, r *http.Request, successType string) {
	query := url.Values{}
	query.Set("success", successType)

	ToHome(w, r, query)
}

func ToTargetOrHome(w http.ResponseWriter, r *http.Request, rawTargetURL string, query url.Values) {
	if rawTargetURL == "" {
		ToHome(w, r, query)
		return
	}

	targetURL, err := url.Parse(rawTargetURL)
	if err != nil {
		ToHome(w, r, query)
		return
	}

	targetQuery := targetURL.Query()
	for key, values := range query {
		for _, value := range values {
			targetQuery.Set(key, value)
		}
	}

	targetURL.RawQuery = targetQuery.Encode()
	http.Redirect(w, r, targetURL.String(), http.StatusSeeOther)
}

func WithSuccess(rawURL string, successType string) string {
	redirectURL, err := url.Parse(rawURL)
	if err != nil {
		return "/?success=" + url.QueryEscape(successType)
	}

	query := redirectURL.Query()
	query.Set("success", successType)
	redirectURL.RawQuery = query.Encode()

	return redirectURL.String()
}
