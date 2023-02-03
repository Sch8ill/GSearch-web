package handlers

import (
	"net/http"

	"github.com/Sch8ill/GSearch/Web/logger"
)

const StaticDir = "webserver/static/"


func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, StaticDir + "index.html")

	logger.LogHTTPRequest(r)
}

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, StaticDir + "search.html")

	logger.LogHTTPRequest(r)
}

func SearchAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")

	query := r.URL.Query().Get("query")

	http.ServeFile(w, r, StaticDir + "test/test_search_query.json")

	logger.LogHTTPRequest(r)
	logger.LogAPIQuery(r, query)
}