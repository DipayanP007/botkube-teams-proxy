package initializers

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		// log.Panic("NO env file")
		log.Panic().Msg("No env file")
	}
}

func LoadENV() {
	initializeLogger()
	log.Info().Msg("Loading environment variables")
}
