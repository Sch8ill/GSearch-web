package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sch8ill/gsearch-web/config"
	"github.com/sch8ill/gsearch-web/db"
	"github.com/sch8ill/gsearch-web/logger"
)

type SearchApiResponse struct {
	Query   string
	Results []db.Site
}

var DB *db.DBClient

func SearchAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")

	query := r.URL.Query().Get("q")
	resStruct := SearchApiResponse{Query: query}

	results := DB.TextSearch(query)

	for _, result := range results {
		result.Text = result.Text[:10]
		resStruct.Results = append(resStruct.Results, result)
	}

	jsonRes, _ := json.Marshal(resStruct)

	w.Write(jsonRes)

	logger.LogAPIQuery(r, query)
}

func VersionAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")

	io.WriteString(w, "GSearch-web@"+config.Version)

	logger.LogAPIQuery(r, "version")
}
