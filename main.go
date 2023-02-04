package main

import (
	"errors"
	"net/http"

	"github.com/Sch8ill/GSearch-web/logger"
	"github.com/Sch8ill/GSearch-web/webserver/handlers"
)



func main() {
	logger.Logger.Info().Msg("Webserver starting...")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/search", handlers.Search)
	http.HandleFunc("/api/v1/search", handlers.SearchAPI)

	http.Handle("/static/css/", http.StripPrefix("/static/css", http.FileServer(http.Dir(handlers.StaticDir + "css/"))))

	err := http.ListenAndServe("localhost:8000", nil)

	// error handling
	if errors.Is(err, http.ErrServerClosed) {
		logger.Logger.Info().Msg("Webserver shut down.")

	} else if err != nil {
		logger.Logger.Error().Err(err).Msg("")
	}
}