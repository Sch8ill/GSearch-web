package logger

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const timestampFormat = "2006-01-02T15:04:05.000"

func CreateLogger(level zerolog.Level) {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(getConsoleWriter()).Level(level).With().Timestamp().Logger()
}

func getConsoleWriter() io.Writer {
	return zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: timestampFormat,
	}
}

func LogHTTPRequest(r *http.Request) {
	log.Debug().
		Str("addr", r.RemoteAddr).
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Msg("")
}

func LogAPIQuery(r *http.Request, query string) {
	log.Info().
		Str("addr", r.RemoteAddr).
		Str("query", query).
		Str("path", r.URL.Path).
		Msg("")
}

func LogError(err error) {
	log.Fatal().Err(err).Msg("")
}
