package link

import "net/http"

type linkForm struct {
	URL   string
	Title string
	Tags  string
}

func parseLinkForm(r *http.Request) linkForm {
	return linkForm{
		URL:   r.FormValue("url"),
		Title: r.FormValue("title"),
		Tags:  r.FormValue("tags"),
	}
}
