package entity

import "net/http"

type RedirectAll struct {
	Target string
	Code   int
}

func (h RedirectAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := h.Target + r.URL.RequestURI()
	http.Redirect(w, r, url, h.Code)
}
