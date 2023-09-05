package handlers

import (
	"net/http"

	"github.com/sch8ill/gsearch-web/logger"
)

func Home(w http.ResponseWriter, r *http.Request) {
	staticHTML("index.html", w, r)
}

func Search(w http.ResponseWriter, r *http.Request) {
	staticHTML("search.html", w, r)
}

func staticHTML(filename string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "static/" + filename)

	logger.LogHTTPRequest(r)
}

func StaticFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])

	logger.LogHTTPRequest(r)
}