package logger

import (
	"os"
	"net/http"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func LogHTTPRequest(r *http.Request) {
	Logger.Debug().
        Str("addr", r.RemoteAddr).
        Str("method", r.Method).
        Str("path", r.URL.Path).
		Msg("")
}

func LogAPIQuery(r *http.Request, query string) {
	Logger.Debug().
        Str("addr", r.RemoteAddr).
        Str("query", query).
		Str("path", r.URL.Path).
		Msg("")
}