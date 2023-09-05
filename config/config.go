package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const Version string = "0.1.0"

const IP string = ""
const Port int = 8000

const LogLevel zerolog.Level = zerolog.DebugLevel

var MongodbURI string = getEnvKey("MONGODBURI")

const DBName string = "gsearch"
const SiteColl string = "sites"

func getEnvKey(key string) string {
	godotenv.Load(".env")
	envVar := os.Getenv(key)
	if envVar == "" {
		log.Logger.Fatal().Msg("Could not load environment variable: " + key)
	}
	return envVar
}
