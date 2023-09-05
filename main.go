package main

import (
	"errors"
	"net/http"
	"runtime"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/sch8ill/gsearch-web/config"
	"github.com/sch8ill/gsearch-web/db"
	"github.com/sch8ill/gsearch-web/handlers"
	"github.com/sch8ill/gsearch-web/logger"
)

func main() {
	logger.CreateLogger(config.LogLevel)
	log.Info().Msg("GSearch-web starting...")
	log.Info().Msg("Golang version:\t " + runtime.Version())
	log.Info().Msg("GSearch version:\t GSearch-web@" + config.Version)
	log.Info().Msg("MongodbURI:\t\t " + config.MongodbURI)

	bindAddr := config.IP + ":" + strconv.FormatInt(int64(config.Port), 10)
	log.Info().Msg("Server addr:\t " + bindAddr)

	registerRoutes()
	
	handlers.DB = db.New(config.MongodbURI)
	handlers.DB.Connect()

	err := http.ListenAndServe(bindAddr, nil)

	// error handling
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Msg("Webserver shut down.")

	} else if err != nil {
		logger.LogError(err)
	}
}

func registerRoutes() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/search", handlers.Search)
	http.HandleFunc("/api/search", handlers.SearchAPI)
	http.HandleFunc("/api/version", handlers.VersionAPI)

	http.HandleFunc("/static/", handlers.StaticFile)
}
