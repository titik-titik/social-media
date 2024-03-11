package container

import (
	"github.com/rs/zerolog"
	"os"
	"social-media/internal/config"
)

func NewLogger(env *config.EnvConfig) *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Level(zerolog.Level(env.Logger.Level))

	return &log
}
