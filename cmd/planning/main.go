package main

import (
	"os"

	"planning"

	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	log.Info().
		Msg("Starting application")

	app := planning.New(
		planning.WithLogger(log),
		planning.OpenBrowser,
	)
	if err := app.Run(); err != nil {
		log.Error().
			Err(err).
			Msg("run")
		os.Exit(1)
	}
}
